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
	Create(userId int, note notebook.Note) (int, error)
	GetAll(userId int) ([]notebook.Note, error)
	GetById(userId, noteId int) (notebook.Note, error)
	Delete(userId, noteId int) error
	Update(userId, noteId int, input notebook.UpdateNoteInput) error
}

type Repository struct {
	Authorization
	Note
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          NewNotePostgres(db),
	}
}
