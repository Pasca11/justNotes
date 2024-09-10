package service

import "github.com/Pasca11/justNotes/internal/repository"

type Service interface {
}

type ServiceImpl struct {
	DB repository.UserRepo
}

func New(db repository.UserRepo) Service {
	return &ServiceImpl{
		DB: db,
	}
}
