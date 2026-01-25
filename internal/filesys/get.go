package filesys

import (
	"io/fs"
	"log"
	"path/filepath"
	"ricer/internal/config"
	"ricer/internal/consts"
	"ricer/internal/types"
)

func GetThemeByName(name string) types.ThemeFile {
	themes := GetThemes()
	for _, theme := range themes {
		if theme.Name != name {
			continue
		}
		return theme
	}
	panic("no theme")
}

func GetCurrentTheme() types.ThemeFile {
	conf := config.GetConfig()
	themes := getFiles(conf.ThemesPath)

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] != consts.CURRENT_THEME_FILE_NAME {
			continue
		}
		return types.ThemeFile{Path: filePath, Name: fileName}
	}

	log.Fatal("Failed to get current theme")
	return types.ThemeFile{}
}

func GetThemes() []types.ThemeFile {
	conf := config.GetConfig()
	themes := getFiles(conf.ThemesPath)
	var themesFiltered []types.ThemeFile

	curThemeFileNameLen := len(consts.CURRENT_THEME_FILE_NAME)
	for _, filePath := range themes {
		fileName := GetFileName(filePath)
		if filePath[len(filePath)-curThemeFileNameLen:] == consts.CURRENT_THEME_FILE_NAME {
			continue
		}
		themesFiltered = append(themesFiltered, types.ThemeFile{Path: filePath, Name: fileName})
	}

	return themesFiltered
}

// files that changing
// todo filter excluded
func getSubjectFiles() []string {
	conf := config.GetConfig()
	filesAll := getFiles(conf.SubjectPath)
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
		log.Fatal(err)
	}
	return files
}

func GetFileName(path string) string {
	return filepath.Base(path)
}
