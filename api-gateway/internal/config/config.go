package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type APIConfig struct{
	Env string `yaml:"env"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustLoad() *APIConfig{
	configPath := os.Getenv("API_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("API_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to load config file: %v\n", err)
	}

	var cfg APIConfig

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil{
		log.Fatalf("failed to read config: %v\n", err)
	}

	return &cfg
}