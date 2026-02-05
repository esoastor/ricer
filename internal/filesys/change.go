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
		content = ReplaceByChangeMap(changeMap, content, subjectPath)
		os.WriteFile(subjectPath, []byte(content), 0644)
	}
	setCurrentTheme(theme)
}

func ReplaceByChangeMap(changeMap []types.ChangeMap, content, contentPath string) string {
	// log.Print(content + "\n")
	for _, change := range changeMap {
		changeFilePathLen := len(change.FilePath)
		contentFilePathLen := len(contentPath)

		if changeFilePathLen > 0 && changeFilePathLen < contentFilePathLen {
			sub := contentPath[contentFilePathLen - changeFilePathLen:]
			if sub != change.FilePath {
				continue
			}
		}
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

