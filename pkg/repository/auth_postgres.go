package repository

import (
	"fmt"
	"github.com/Smolvika/notebook.git"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(bd *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: bd}
}

func (p *AuthPostgres) CreateUser(user notebook.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := p.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (p *AuthPostgres) GetUser(username, password string) (notebook.User, error) {
	var user notebook.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := p.db.Get(&user, query, username, password)
	return user, err
}
