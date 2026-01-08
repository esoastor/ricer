package config

import (
	"os"
	"io"
	"log"
	"path/filepath"
	"gopkg.in/yaml.v3"
	"ricer/internal/types"
	"ricer/internal/consts"
)




func GetConfig() types.Config {
	path := getConfigPath()
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config types.Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

	
	return config
}

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return filepath.Join(home, consts.CONFIG_PATH)
}

