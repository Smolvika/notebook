package service

import (
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Note interface {
	Create(userId int, note notebook.Note) (int, error)
	GetAll(userId int) ([]notebook.Note, error)
	GetById(userId, noteId int) (notebook.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input notebook.UpdateNoteInput) error
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNoteService(repos.Note),
	}
}
