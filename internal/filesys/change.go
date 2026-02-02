package filesys

import (
	"path/filepath"
	//"io/fs"
	"log"
	"os"
	"ricer/internal/config"
	"ricer/internal/theme"
	"ricer/internal/consts"
	"ricer/internal/types"
	"strings"
)

func SubmitTheme(theme theme.ThemeFile) {
	subjects := GetSubjectFiles()

	changeMap := CreateChangeMapForCurrent(theme)
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

func setCurrentTheme(theme theme.ThemeFile) {
	conf := config.GetConfig()
	defPath := filepath.Join(conf.ThemesPath, consts.CURRENT_THEME_FILE_NAME)

	content, err := os.ReadFile(theme.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	os.WriteFile(defPath, content, 0644)
}

