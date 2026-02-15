package theme

import (
	"fmt"
	"log"
	"regexp"
	"ricer/internal/types"
	"slices"
	"strings"
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


func processRows(rows []string) (map[string]types.ThemeRow, error) {
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

		skipStep, err := processMultilineValue(rowClear, &meta, &multilineParams, &theme)
		if err != nil {
			return theme, err
		}
		if skipStep {
			continue
		}

		processMeta(row, &meta)
		cleanRowFinally(rowClear)

		first, second, found := strings.Cut(rowClear, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			continue
		}

		err = clearValuesAndAddToTheme(first, second, &meta, &theme)
		if err != nil {
			return theme, err
		}
	}
	return theme, nil
}

func processMultilineValue(row string, meta *types.ThemeRowMeta, multilineParams *types.MultilineStruct, theme *map[string]types.ThemeRow) (bool, error) {
	hasMultilineMark := strings.Contains(row, multilineValueMark)

	if !hasMultilineMark && multilineParams.IsMultilineValue {
		multilineParams.MultilineValue += "\n" + row
		return true, nil
	}

	if hasMultilineMark && !multilineParams.IsMultilineValue {
		multilineParams.IsMultilineValue = true
		first, second, found := strings.Cut(row, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			return true, nil
		}
		key := strings.TrimSpace(first)
		val := strings.TrimSpace(second)
		multilineParams.MultilineKey = key
		multilineParams.MultilineValue = strings.ReplaceAll(val, multilineValueMark, "")
		return true, nil
	}

	if hasMultilineMark && multilineParams.IsMultilineValue {
		multilineParams.IsMultilineValue = false
		multilineParams.MultilineValue += "\n" + strings.ReplaceAll(row, multilineValueMark, "")

		err := clearValuesAndAddToTheme(multilineParams.MultilineKey, multilineParams.MultilineValue, meta, theme)

		if err != nil {
			return false, err
		}

		multilineParams.MultilineKey = ""
		multilineParams.MultilineValue = ""

		return true, nil
	}

	return false, nil
}

func processMeta(row string, meta *types.ThemeRowMeta) error {
	const failReadPanicMessage = "Invalid theme structure, failed to read row:\n"
	rowLTrim := strings.TrimLeft(row, " ")
	if len(rowLTrim) == 0 {
		return nil
	}
	isTag := string(rowLTrim[0]) == metaOpen

	if !isTag {
		return nil
	}

	metaPathLen := len(meta.Path)
	tag, value := ParseStartMeta(rowLTrim)
	if tag != "" && value != "" {
		if metaPathLen > 0 {
			return fmt.Errorf("%s%s", failReadPanicMessage, row)
		}

		meta.Path = value
		return nil
	}

	tag = ParseEndMeta(rowLTrim)

	if tag == "" || metaPathLen == 0 {
		return fmt.Errorf("%s%s", failReadPanicMessage, row)
	}

	meta.Path = ""
	return nil
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

func clearValuesAndAddToTheme(keyRaw, valRaw string, meta *types.ThemeRowMeta, theme *map[string]types.ThemeRow) error {
	key := strings.TrimSpace(keyRaw)
	val := strings.TrimSpace(valRaw)

	mapKey := key + meta.Path
	_, exists := (*theme)[mapKey]
	if exists {
		return fmt.Errorf("Invalid theme syntax: key repeat forbidden (%s)", mapKey)
	}
	(*theme)[mapKey] = types.ThemeRow{
		Key:   key,
		Value: val,
		Meta:  *meta,
	}
	return nil
}
