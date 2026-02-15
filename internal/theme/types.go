package theme

import (
	"fmt"
	"log"
	"os"
	"ricer/internal/types"
	"strings"
)

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

	theme, err := processRows(rows)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return theme
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
