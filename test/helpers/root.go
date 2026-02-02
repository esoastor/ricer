package helpers

import (
	"ricer/internal/theme"
	"os"
	"path/filepath"
)

// test helpers
const testThemesPath = "test/data/"
const goodThemeFileName = "theme-good"
const badThemeFileName = "theme-bad"
const diffThemeFileName = "theme-good-diff"

func GetBadTheme() theme.ThemeFile {
	return getThemeFile(badThemeFileName)
}
func GetGoodTheme() theme.ThemeFile {
	return getThemeFile(goodThemeFileName)
}

func GetDiffTheme() theme.ThemeFile {
	return getThemeFile(diffThemeFileName)
}

func getThemeFile(themeFileName string) theme.ThemeFile {
	projectPath, err := os.Getwd()
	if err != nil {
		panic("Filed to resolve project path")
	}
	path := filepath.Join(projectPath, "../..", testThemesPath, themeFileName)
	file := theme.ThemeFile{}
	file.Path = path
	file.Name = "test"
	return file
}
