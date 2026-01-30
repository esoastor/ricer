package filesys

import (
	"path/filepath"
	//"io/fs"
	"log"
	"os"
	"ricer/internal/config"
	"ricer/internal/consts"
	"ricer/internal/types"
	"strings"
)

func SubmitTheme(theme types.ThemeFile) {
	subjects := GetSubjectFiles()

	changeMap := CreateChangeMap(theme)
	for _, subjectPath := range subjects {
		contentRaw, err := os.ReadFile(subjectPath)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		content := string(contentRaw)
		content = replaceByTheme(changeMap, content)
		os.WriteFile(subjectPath, []byte(content), 0644)
	}
	setCurrentTheme(theme)
}

func replaceByTheme(changeMap []types.ChangeMap, content string) string {
	for _, change := range changeMap {
		content = strings.ReplaceAll(content, change.From, change.To)
	}
	return content
}

func setCurrentTheme(theme types.ThemeFile) {
	conf := config.GetConfig()
	defPath := filepath.Join(conf.ThemesPath, consts.CURRENT_THEME_FILE_NAME)

	content, err := os.ReadFile(theme.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	os.WriteFile(defPath, content, 0644)
}

func CreateChangeMap(themeFile types.ThemeFile) []types.ChangeMap {
	current := GetCurrentTheme()

	newTheme := themeFile.FormTheme()
	curTheme := current.FormTheme()

	var changeMap []types.ChangeMap
	for code, value := range curTheme {
		if newTheme[code] == "" || value == newTheme[code] {
			continue
		}
		changeMap = append(changeMap, types.ChangeMap{
			Code:  code,
			From: value,
			To:   newTheme[code],
		})
	}
	return changeMap
}

