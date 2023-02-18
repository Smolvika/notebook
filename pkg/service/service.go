package service

import "github.com/Smolvika/notebook.git/pkg/repository"

type Authorization interface {
}

type Note interface {
}

type Service struct {
	Authorization
	Note
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
