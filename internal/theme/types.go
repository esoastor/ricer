package theme

import (
	"fmt"
	"os"
	"ricer/internal/types"
	"strings"
)

type ThemeFile struct {
	types.ThemeFile
}

func (tf ThemeFile) FormTheme() (map[string]types.ThemeRow, error) {
	contentRaw, errRead := os.ReadFile(tf.Path)
	if errRead != nil {
		emptyTheme := make(map[string]types.ThemeRow)
		return emptyTheme, errRead
	}
	content := string(contentRaw)
	rows := strings.Split(content, "\n")

	theme, errProcess := processRows(rows)
	return theme, errProcess
}

type ThemeFileCollection struct {
	Files []ThemeFile
}

func (tfc ThemeFileCollection) GetByName(name string) ThemeFile {
	for _, file := range tfc.Files {
		if file.Name != name {
			continue
		}
		return file
	}
	fmt.Printf("Theme not found: \"%s\"\n", name)
	os.Exit(1)
	return ThemeFile{}
}
