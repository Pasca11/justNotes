package service

import (
	authv1 "github.com/Pasca11/gRPC-Auth/proto/gen"
	"github.com/Pasca11/justNotes/models"
)

type UserService1 interface {
	Login(req *authv1.LoginRequest) (*authv1.LoginResponse, error)
	Register(req *authv1.RegisterRequest) (*authv1.RegisterResponse, error)
}

type UserService interface {
	Login(req *models.User) (*models.LoginResponse, error)
	Register(req *models.User) (*models.User, error)
}

type NotesService interface {
	GetNotes(id int) ([]models.Note, error)
	CreateNote(id int, note *models.Note) error
	DeleteNote(id int) error
}
