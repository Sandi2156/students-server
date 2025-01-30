package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags
	}

	if configPath == "" {
		log.Fatal("Config file is required (NOT FOUND).")
	}

	if _, err := os.Stat(configPath); os.IsNotExist((err)) {
		log.Fatalf("Config file path is wrong %v", configPath)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("Config file can not be parsed: %v", err.Error())
	}

	return &config
}