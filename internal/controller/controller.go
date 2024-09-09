package controller

import "github.com/Pasca11/justNotes/internal/service"

type Controller interface {
}

type ControllerImpl struct {
	service service.Service
}

func New(cfg *Config) (Controller, error) {
	return nil, nil
}
