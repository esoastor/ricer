package theme

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"ricer/internal/types"
)

const commentMark = "//"
const multilineValueMark = "```"
const keyValSeparator = "="

// meta
const metaOpen = "["
const metaClose = "]"

var metaAllowedTags = []string{"file"}

const metaStartRegex = `^\s*\` + metaOpen + `\s*(\w+)\s+([\w|/|.]+)\s*\` + metaClose + `\s*$`
const metaEndRegex = `^\s*\` + metaOpen + `\s*end\s*(\w+)\s*\` + metaClose + `\s*$`


type ThemeFile struct {
	types.ThemeFile
}

func (tf ThemeFile) FormTheme() map[string]types.ThemeRow {
	return GetTheme(tf)
}

func GetTheme(tf ThemeFile) map[string]types.ThemeRow {
	contentRaw, err := os.ReadFile(tf.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(contentRaw)
	rows := strings.Split(content, "\n")

	theme := processRows(rows)

	return theme
}

func processRows(rows []string) map[string]types.ThemeRow {
	theme := make(map[string]types.ThemeRow)

	meta := types.ThemeRowMeta{
		Path: "",
	}

	multilineParams := types.MultilineStruct{
		IsMultilineValue: false,
		MultilineKey:     "",
		MultilineValue:   "",
	}

	for _, row := range rows {
		rowClear := cleanRowInitialy(row)

		skipStep := processMultilineValue(rowClear, &meta, &multilineParams, &theme)
		if skipStep {
			continue
		}

		processMeta(row, &meta)
		cleanRowFinally(rowClear)

		first, second, found := strings.Cut(rowClear, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			continue
		}

		clearValuesAndAddToTheme(first, second, &meta, &theme)
	}
	return theme
}

func processMultilineValue(row string, meta *types.ThemeRowMeta, multilineParams *types.MultilineStruct, theme *map[string]types.ThemeRow) bool {
	hasMultilineMark := strings.Contains(row, multilineValueMark)

	if !hasMultilineMark && multilineParams.IsMultilineValue {
		multilineParams.MultilineValue += "\n" + row
		return true
	}

	if hasMultilineMark && !multilineParams.IsMultilineValue {
		multilineParams.IsMultilineValue = true
		first, second, found := strings.Cut(row, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			return true
		}
		key := strings.TrimSpace(first)
		val := strings.TrimSpace(second)
		multilineParams.MultilineKey = key
		multilineParams.MultilineValue = strings.ReplaceAll(val, multilineValueMark, "")
		return true
	}

	if hasMultilineMark && multilineParams.IsMultilineValue {
		multilineParams.IsMultilineValue = false
		multilineParams.MultilineValue += "\n" + strings.ReplaceAll(row, multilineValueMark, "")

		clearValuesAndAddToTheme(multilineParams.MultilineKey, multilineParams.MultilineValue, meta, theme)

		multilineParams.MultilineKey = ""
		multilineParams.MultilineValue = ""

		return true
	}

	return false
}

func processMeta(row string, meta *types.ThemeRowMeta) {
	const failReadPanicMessage = "invalid theme structure, failed to read"
	rowLTrim := strings.TrimLeft(row, " ")
	if len(rowLTrim) == 0 {
		return
	}
	isTag := string(rowLTrim[0]) == metaOpen

	if !isTag {
		return
	}

	metaPathLen := len(meta.Path)
	tag, value := ParseStartMeta(rowLTrim)
	if tag != "" && value != "" {
		if metaPathLen > 0 {
			panic(failReadPanicMessage)
		}

		meta.Path = value
		return
	}

	tag = ParseEndMeta(rowLTrim)

	if tag == "" || metaPathLen == 0  {
		panic(failReadPanicMessage)
	}


	meta.Path = ""	
}

func ParseStartMeta(row string) (string, string) {
	re := regexp.MustCompile(metaStartRegex)
	found := re.FindStringSubmatch(row)
	isLenOk := len(found) == 3

	if !isLenOk {
		return "", ""
	}

	metaTag := found[1]
	metaValue := found[2]

	if !slices.Contains(metaAllowedTags, metaTag) {
		log.Printf("ivalid tag")
		return "", ""
	}

	return metaTag, metaValue
}

func ParseEndMeta(row string) string {
	re := regexp.MustCompile(metaEndRegex)
	found := re.FindStringSubmatch(row)
	isLenOk := len(found) == 2

	if !isLenOk {
		return ""
	}

	metaTag := found[1]

	if !slices.Contains(metaAllowedTags, metaTag) {
		log.Printf("ivalid tag")
		return ""
	}

	return metaTag
}
func cleanRowInitialy(row string) string {
	regComment := regexp.MustCompile(commentMark + `.*`)
	rowClear := regComment.ReplaceAllString(row, "")
	return rowClear
}

func cleanRowFinally(row string) string {
	row = strings.TrimSpace(row)
	reSp := regexp.MustCompile(`\s+`)
	return reSp.ReplaceAllString(row, " ")
}

func clearValuesAndAddToTheme(keyRaw, valRaw string, meta *types.ThemeRowMeta, theme *map[string]types.ThemeRow) {
	key := strings.TrimSpace(keyRaw)
	val := strings.TrimSpace(valRaw)

	mapKey := key + meta.Path
	_, exists := (*theme)[mapKey]
	if (exists) {
		panic("Invalid theme (key repeat forbidden) " + mapKey)
	}
	(*theme)[mapKey] = types.ThemeRow{
		Key:   key,
		Value: val,
		Meta:  *meta,
	}
}
