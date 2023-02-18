package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Note interface {
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
