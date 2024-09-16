package config

import (
	"github.com/Pasca11/justNotes/internal/logger"
	"github.com/Pasca11/justNotes/internal/transport/server"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	App    *App           `yaml:"app"`
	Server *server.Config `yaml:"server"`
	Logger *logger.Config `yaml:"logger"`
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
