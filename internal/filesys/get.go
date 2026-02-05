package filesys

import (
	"io/fs"
	"log"
	"path/filepath"
	"ricer/internal/config"
	"ricer/internal/consts"
	"ricer/internal/theme"
	"ricer/internal/types"
)

func GetThemeByName(name string) theme.ThemeFile {
	themes := GetThemes()
	for _, theme := range themes {
		if theme.Name != name {
			continue
		}
		return theme
	}
	panic("no theme")
}

func GetCurrentTheme() theme.ThemeFile {
	conf := config.GetConfig()
	themes := GetFiles(conf.ThemesPath)

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] != consts.CURRENT_THEME_FILE_NAME {
			continue
		}

		file := theme.ThemeFile{}
		file.Path = filePath
		file.Name = fileName
		return file 
	}

	log.Fatal("Failed to get current theme")
	return theme.ThemeFile{}
}

func GetThemes() []theme.ThemeFile {
	conf := config.GetConfig()
	themes := GetFiles(conf.ThemesPath)
	var themesFiltered []theme.ThemeFile

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] == consts.CURRENT_THEME_FILE_NAME {
			continue
		}
		file := theme.ThemeFile{}
		file.Path = filePath
		file.Name = fileName
		themesFiltered = append(themesFiltered, file)
	}

	return themesFiltered
}

func GetSubjectFiles() []string {
	conf := config.GetConfig()
	filesAll := GetFiles(conf.SubjectPath)
	var filesFiltered []string
	excludes := conf.Exclude
	for _, file := range filesAll {
		write := true
		for _, exclude := range excludes {
			exLen := len(exclude)
			if exLen > len(file) {
				filesFiltered = append(filesFiltered, file)
				continue
			}
			fileSub := file[0:exLen]

			if fileSub == exclude {
				write = false
				break
			}
		}
		if write {
			filesFiltered = append(filesFiltered, file)
		}
	}
	return filesFiltered
}

func GetFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.WalkDir(path, func(currentPath string, directoryEntry fs.DirEntry, walkError error) error {
		if walkError != nil {
			return walkError
		}

		if directoryEntry.IsDir() {
			return nil
		}
		files = append(files, currentPath)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetFileName(path string) string {
	return filepath.Base(path)
}

func CreateChangeMapForCurrent(themeFile theme.ThemeFile) []types.ChangeMap {
	current := GetCurrentTheme()

	return CreateChangeMap(current, themeFile)
}


func CreateChangeMap(from, to theme.ThemeFile) []types.ChangeMap {
	newTheme := to.FormTheme()
	curTheme := from.FormTheme()
    
	var changeMap []types.ChangeMap
	for _, row := range curTheme {
		pairNewThemeRow, exists := newTheme[row.FormId()]
		if !exists || row.Value == pairNewThemeRow.Value {
			continue
		}

		changeMap = append(changeMap, types.ChangeMap{
			Code:  row.Key,
			From: row.Value,
			To:   pairNewThemeRow.Value,
			FilePath: row.Meta.Path,
		})
	}
	return changeMap
}

