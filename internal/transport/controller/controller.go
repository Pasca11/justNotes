package controller

import (
	"encoding/json"
	authv1 "github.com/Pasca11/gRPC-Auth/proto/gen"
	"github.com/Pasca11/justNotes/internal/metrics"
	"github.com/Pasca11/justNotes/internal/service"
	"github.com/Pasca11/justNotes/models"
	"github.com/Pasca11/justNotes/pkg/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	_ "github.com/Pasca11/justNotes/docs"
)

type Controller interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetNotes(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	notesService service.NotesService
	userService  service.UserService
	log          logger.Logger
}

func New(s service.NotesService, s2 service.UserService, l logger.Logger) (Controller, error) {
	return &ControllerImpl{
		notesService: s,
		userService:  s2,
		log:          l,
	}, nil
}

// Login handles login requests
// @summary Authenticate user
// @tags Auth
// @description enter credentials to login
// @accept json
// @produce json
// @param credentials body models.User true "username and password"
// @success 200
// @Failure 500
// @failure 400
// @router /auth/login [post]
func (c *ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	metrics.LoginCount.Inc()

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := c.userService.Login(&authv1.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
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

// Register handles register requests
// @summary Create user
// @tags Auth
// @description enter credentials to register
// @accept json
// @param credentials body models.User true "username and password"
// @success 200
// @Failure 500
// @failure 400
// @router /auth/register [post]
func (c *ControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	metrics.RegisterCounter.Inc()

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Role = "user"
	_, err = c.userService.Register(&authv1.RegisterRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

// GetNotes return notes
// @summary Get user`s notes
// @tags notes
// @description returns all user`s notes
// @accept json
// @produce json
// @success 200
// @Failure 500
// @failure 401
// @router /auth/notes [get]
func (c *ControllerImpl) GetNotes(w http.ResponseWriter, r *http.Request) {
	//token := r.Header.Get("Authorization")
	//userId, err := notesService.ExtractUserIdFromToken(token)
	//if err != nil {
	//	c.log.Error(err.Error())
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	notes, err := c.notesService.GetNotes(userId)
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

// CreateNote handles register requests
// @summary Create note
// @tags notes
// @description enter text and deadline(optional) to create note
// @accept json
// @param note body models.Note true "text and deadline"
// @success 200
// @failure 500
// @failure 401
// @router /auth/notes [post]
func (c *ControllerImpl) CreateNote(w http.ResponseWriter, r *http.Request) {
	note := &models.Note{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//token := r.Header.Get("Authorization")
	//userId, err := notesService.ExtractUserIdFromToken(token)
	//if err != nil {
	//	c.log.Error(err.Error())
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.notesService.CreateNote(userId, note)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteNote handles register requests
// @summary Delete note
// @tags notes
// @description enter id to delete note
// @accept json
// @param note body models.Note true "note id"
// @success 200
// @failure 500
// @router /auth/notes [delete]
func (c *ControllerImpl) DeleteNote(w http.ResponseWriter, r *http.Request) {
	delNote := &models.Note{}
	err := json.NewDecoder(r.Body).Decode(delNote)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.notesService.DeleteNote(delNote.ID)
	if err != nil {
		c.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
