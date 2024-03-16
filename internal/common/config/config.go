package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type config struct {
	Http struct {
		Port int16 `yaml:"port"`
	} `yaml:"http"`
	DatabaseURL string `yaml:"databaseURL"`
}

var (
	cfg  config
	once sync.Once
)

func HttpPort() int16 {
	return cfg.Http.Port
}

func DatabaseURL() string {
	return cfg.DatabaseURL
}

func MustReadConfigFromFile(filepath string) {
	once.Do(func() {
		data, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatalf("Reading config from file %s: %v\n", filepath, err)
		}

		err = json.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatalf("Unmarshalling config from file %s: %v\n", filepath, err)
		}
	})

	return
}
