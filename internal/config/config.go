package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	App *App `yaml:"app"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	err = cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
