package controller

import (
	"encoding/json"
	"github.com/Pasca11/justNotes/internal/logger"
	"github.com/Pasca11/justNotes/internal/service"
	"github.com/Pasca11/justNotes/models"
	"net/http"
)

type Controller interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetNotes(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	service service.UserService
	log     logger.Logger
}

func New(s service.UserService, l logger.Logger) (Controller, error) {
	return &ControllerImpl{
		service: s,
		log:     l,
	}, nil
}

func (c *ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := c.service.Login(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.service.Register(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ControllerImpl) GetNotes(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userId, err := service.ExtractUserIdFromToken(token)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	notes, err := c.service.GetNotes(userId)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ControllerImpl) CreateNote(w http.ResponseWriter, r *http.Request) {
	note := &models.Note{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	userId, err := service.ExtractUserIdFromToken(token)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = c.service.CreateNote(userId, note)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ControllerImpl) DeleteNote(w http.ResponseWriter, r *http.Request) {
	delNote := &models.Note{}
	err := json.NewDecoder(r.Body).Decode(delNote)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.service.DeleteNote(delNote.ID)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
