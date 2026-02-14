package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"ricer/internal/consts"
	"ricer/internal/types"
)

func Get() types.Config { // todo just Get()
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
