package filesys

import (
	"io/fs"
	"log"
	"path/filepath"
	"ricer/internal/config"
)

func GetSubjectFiles() []string {
	conf := config.Get()
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

