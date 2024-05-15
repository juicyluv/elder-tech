package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type config struct {
	Http struct {
		Port          int16  `yaml:"port"`
		HTTPSKeyPath  string `yaml:"https_key_path"`
		HTTPSCertPath string `yaml:"https_cert_path"`
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

func HttpsCertPath() string {
	return cfg.Http.HTTPSCertPath
}

func HttpsKeyPath() string {
	return cfg.Http.HTTPSKeyPath
}

func MustReadConfigFromFile(filepath string) {
	once.Do(func() {
		data, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatalf("Reading config from file %s: %v\n", filepath, err)
		}

		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatalf("Unmarshalling config from file %s: %v\n", filepath, err)
		}
	})
}
