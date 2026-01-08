package filesys

import (
	"path/filepath"
	"io/fs"
	"ricer/internal/config"
	"ricer/internal/consts"
)

func GetChangingConfigFiles() []string {
	conf := config.GetConfig()
	return getFiles(conf.SubjectPath)
}

func GetThemes() []string {
	conf := config.GetConfig()
	themes := getFiles(conf.ThemesPath)
	var themesFiltered []string 

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, file := range themes {
		if file[len(file)-curThemeFileNameLen:] == consts.CURRENT_THEME_FILE_NAME {
			continue
		}
		themesFiltered = append(themesFiltered, file)
	}

	return themesFiltered
}

func getFiles(path string) []string {
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
		panic(err)
	}
	return files
}

