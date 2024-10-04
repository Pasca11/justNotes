package repository

import "github.com/Pasca11/justNotes/models"

type UserRepo interface {
	//CreateUser(user *models.User) error
	//GetUser(username string) (*models.User, error)
	GetNotes(user_id int) ([]models.Note, error)
	CreateNote(id int, note *models.Note) error
	DeleteNote(id int) error
}
