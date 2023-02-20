package repository

import (
	"github.com/Smolvika/notebook.git"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GetUser(username, password string) (notebook.User, error)
}

type Note interface {
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          nil,
	}
}
