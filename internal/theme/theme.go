package theme

import (
	"log"
	"path/filepath"
	"os"
	"ricer/internal/config"
	"ricer/internal/types"
	"ricer/internal/consts"
	"ricer/internal/filesys"
	"strings"
)

func GetCurrent() ThemeFile {
	conf := config.Get()
	themes := filesys.GetFiles(conf.ThemesPath)

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := filesys.GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] != consts.CURRENT_THEME_FILE_NAME {
			continue
		}

		file := ThemeFile{}
		file.Path = filePath
		file.Name = fileName
		return file
	}

	log.Fatal("Failed to get current theme")
	return ThemeFile{}
}

// get all awailable themes
func GetAll() ThemeFileCollection {
	conf := config.Get()
	themes := filesys.GetFiles(conf.ThemesPath)
	var themesFiltered []ThemeFile

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := filesys.GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] == consts.CURRENT_THEME_FILE_NAME {
			continue
		}
		file := ThemeFile{}
		file.Path = filePath
		file.Name = fileName
		themesFiltered = append(themesFiltered, file)
	}

	return ThemeFileCollection{Files: themesFiltered}
}

func CreateChangeMapForCurrent(themeFile ThemeFile) []types.ChangeMap {
	current := GetCurrent()

	return CreateChangeMap(current, themeFile)
}

func CreateChangeMap(from, to ThemeFile) []types.ChangeMap {
	newTheme := to.FormTheme()
	curTheme := from.FormTheme()

	var changeMap []types.ChangeMap
	for _, row := range curTheme {
		pairNewThemeRow, exists := newTheme[row.FormId()]
		if !exists || row.Value == pairNewThemeRow.Value {
			continue
		}

		changeMap = append(changeMap, types.ChangeMap{
			Code:     row.Key,
			From:     row.Value,
			To:       pairNewThemeRow.Value,
			FilePath: row.Meta.Path,
		})
	}
	return changeMap
}

// set theme from file as current
func Submit(theme ThemeFile) {
	subjects := filesys.GetSubjectFiles()

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
	setCurrent(theme)
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

// write current theme to 'current' file
func setCurrent(theme ThemeFile) {
	conf := config.Get()
	defPath := filepath.Join(conf.ThemesPath, consts.CURRENT_THEME_FILE_NAME)

	content, err := os.ReadFile(theme.Path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	os.WriteFile(defPath, content, 0644)
}

