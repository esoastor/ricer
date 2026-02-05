package helpers

import (
	"log"
	"os"
	"path/filepath"
	"ricer/internal/filesys"
	"ricer/internal/theme"
)

type TestFile struct {
	Path string
	Content string
}

const testThemesPath = "test/data/themes"
const subjectsPath = "test/data/subjects"
const expectedPath = "test/data/expected"

const goodThemeFileName = "theme-good"
const badThemeFileName = "theme-bad"
const diffThemeFileName = "theme-good-diff"
const subjectCurrentFileName = "subjects-current"
const subjectThemeFileName = "subjects-theme"

func GetSubjects() map[string]TestFile {
	return getFilesList(subjectsPath)
}

func GetExpected() map[string]TestFile {
	return getFilesList(expectedPath)
}

func GetBadTheme() theme.ThemeFile {
	return getThemeFile(badThemeFileName)
}

func GetGoodTheme() theme.ThemeFile {
	return getThemeFile(goodThemeFileName)
}

func GetDiffTheme() theme.ThemeFile {
	return getThemeFile(diffThemeFileName)
}

func GetSubjectCurrentTheme() theme.ThemeFile {
	return getThemeFile(subjectCurrentFileName)
}

func GetSubjectThemeTheme() theme.ThemeFile {
	return getThemeFile(subjectThemeFileName)
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

func getFilesList(dir string) map[string]TestFile {
	projectPath, err := os.Getwd()
	if err != nil {
		panic("Filed to resolve project path")
	}
	path := filepath.Join(projectPath, "../..", dir)

	files := filesys.GetFiles(path)
	subjects := make(map[string]TestFile)

	for _, file := range files {
		contentRaw, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		content := string(contentRaw)
		subjects[filepath.Base(file)] = TestFile{Content: content, Path: file}
	}
	return subjects
}
