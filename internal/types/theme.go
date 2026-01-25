package types

import (
	"log"
	"os"
	"strings"
)

type ThemeFile struct {
	Path string
	Name string
}

type Theme map[string]string

func (tf *ThemeFile) FormTheme() Theme {
	contentRaw, err := os.ReadFile(tf.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(contentRaw)
	rows := strings.Split(content, "\n")

	theme := make(map[string]string)

	for _, row := range rows {

		elems := strings.Split(row, " ")
		if len(elems) != 2 {
			continue
		}
		theme[elems[0]] = elems[1]
	}
	return theme
}
