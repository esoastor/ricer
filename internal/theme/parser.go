package theme

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

type multilineStruct struct {
	isMultilineValue bool
	multilineKey     string
	multilineValue   string
}

const commentMark = "//"
const multilineValueMark = "```"
const keyValSeparator = "="

// meta
const metaOpen = "["
const metaClose = "]"

var metaAllowedTags = []string{"file"}

const metaStartRegex = `^\s*\` + metaOpen + `\s*(\w+)\s+([\w|/|.]+)\s*\` + metaClose + `\s*$`
const metaEndRegex = `^\s*\` + metaOpen + `\s*end\s*(\w+)\s*\` + metaClose + `\s*$`

func GetTheme(tf ThemeFile) []ThemeRow {
	contentRaw, err := os.ReadFile(tf.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(contentRaw)
	rows := strings.Split(content, "\n")

	theme := processRows(rows)

	return theme
}

func processRows(rows []string) []ThemeRow {
	var theme []ThemeRow

	meta := ThemeRowMeta{
		Path: "",
	}

	multilineParams := multilineStruct{
		isMultilineValue: false,
		multilineKey:     "",
		multilineValue:   "",
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

func processMultilineValue(row string, meta *ThemeRowMeta, multilineParams *multilineStruct, theme *[]ThemeRow) bool {
	hasMultilineMark := strings.Contains(row, multilineValueMark)

	if !hasMultilineMark && multilineParams.isMultilineValue {
		multilineParams.multilineValue += "\n" + row
		return true
	}

	if hasMultilineMark && !multilineParams.isMultilineValue {
		multilineParams.isMultilineValue = true
		first, second, found := strings.Cut(row, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			return true
		}
		key := strings.TrimSpace(first)
		val := strings.TrimSpace(second)
		multilineParams.multilineKey = key
		multilineParams.multilineValue = strings.ReplaceAll(val, multilineValueMark, "")
		return true
	}

	if hasMultilineMark && multilineParams.isMultilineValue {
		multilineParams.isMultilineValue = false
		multilineParams.multilineValue += "\n" + strings.ReplaceAll(row, multilineValueMark, "")

		clearValuesAndAddToTheme(multilineParams.multilineKey, multilineParams.multilineValue, meta, theme)

		multilineParams.multilineKey = ""
		multilineParams.multilineValue = ""

		return true
	}

	return false
}

func processMeta(row string, meta *ThemeRowMeta) {
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
	tag, value := parseStartMeta(rowLTrim)
	if tag != "" && value != "" {
		if metaPathLen > 0 {
			panic(failReadPanicMessage)
		}

		meta.Path = value
		return
	}

	tag = parseEndMeta(rowLTrim)

	if tag == "" || metaPathLen == 0  {
		panic(failReadPanicMessage)
	}


	meta.Path = ""	
}

func parseStartMeta(row string) (string, string) {
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

func parseEndMeta(row string) string {
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

func clearValuesAndAddToTheme(keyRaw, valRaw string, meta *ThemeRowMeta, theme *[]ThemeRow) {
	key := strings.TrimSpace(keyRaw)
	val := strings.TrimSpace(valRaw)
	*theme = append(*theme, ThemeRow{
		Key:   key,
		Value: val,
		Meta:  *meta,
	})
}
