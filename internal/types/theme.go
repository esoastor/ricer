package types

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type ThemeFile struct {
	Path string
	Name string
}

type Theme map[string]string

const commentMark = "//"
const multilineValueMark = "```"
const keyValSeparator = "="

func (tf *ThemeFile) FormTheme() Theme {
	contentRaw, err := os.ReadFile(tf.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(contentRaw)
	rows := strings.Split(content, "\n")

	theme := make(map[string]string)

	multilineValue := false
	multilineKey := ""
	multilineValueDummy := ""

	for _, row := range rows {
		reCom := regexp.MustCompile(commentMark + `.*`)
		rowClear := reCom.ReplaceAllString(row, "")

		hasMultilineMark := strings.Contains(rowClear, multilineValueMark)

		if !hasMultilineMark && multilineValue {
			multilineValueDummy += "\n" + rowClear
			continue
		}

		if hasMultilineMark && !multilineValue {
			multilineValue = true
			first, second, found := strings.Cut(rowClear, keyValSeparator)
			if !found || len(first) == 0 || len(second) == 0 {
				continue
			}
			key := strings.TrimSpace(first)
			val := strings.TrimSpace(second)
			multilineKey = key
			multilineValueDummy = strings.ReplaceAll(val, multilineValueMark, "")
			continue
		}

		if hasMultilineMark && multilineValue {
			multilineValue = false
			multilineValueDummy += "\n" + strings.ReplaceAll(rowClear, multilineValueMark, "")

			key := strings.TrimSpace(multilineKey)
			val := strings.TrimSpace(multilineValueDummy)
			theme[key] = val

			multilineKey = ""
			multilineValueDummy = ""

			continue
		}

		rowClear = strings.TrimSpace(rowClear)

		reSp := regexp.MustCompile(`\s+`)
		rowClear = reSp.ReplaceAllString(rowClear, " ")

		first, second, found := strings.Cut(rowClear, keyValSeparator)
		if !found || len(first) == 0 || len(second) == 0 {
			continue
		}
		key := strings.TrimSpace(first)
		val := strings.TrimSpace(second)
		theme[key] = val
	}
	return theme
}
