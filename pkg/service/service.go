package service

import (
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Note interface {
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
