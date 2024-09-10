package controller

import (
	"github.com/Pasca11/justNotes/internal/service"
	"net/http"
)

type Controller interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	service service.Service
}

func New(s service.Service) (Controller, error) {
	return &ControllerImpl{service: s}, nil
}

func (c *ControllerImpl) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
