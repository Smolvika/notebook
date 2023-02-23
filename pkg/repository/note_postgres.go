package repository

import (
	"fmt"
	"github.com/Smolvika/notebook.git"
	"github.com/jmoiron/sqlx"
	"strings"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (p *NotePostgres) Create(userId int, note notebook.Note) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	creatNoteQuery := fmt.Sprintf("INSERT INTO %s (date, description) VALUES ($1, $2) RETURNING id", notesTable)
	row := tx.QueryRow(creatNoteQuery, note.Date, note.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	creatUsersNoteQuery := fmt.Sprintf("INSERT INTO %s(user_id, note_id) VALUES ($1,$2) ", usersNotesTable)
	_, err = tx.Exec(creatUsersNoteQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (p *NotePostgres) GetAll(userId int) ([]notebook.Note, error) {
	var notes []notebook.Note
	query := fmt.Sprintf("SELECT u.id, u.date,u.description FROM %s u JOIN %s us ON u.id = us.note_id WHERE us.user_id = $1", notesTable, usersNotesTable)
	err := p.db.Select(&notes, query, userId)
	return notes, err
}
func (p *NotePostgres) GetById(userId, noteId int) (notebook.Note, error) {
	var note notebook.Note
	query := fmt.Sprintf("SELECT u.id, u.date,u.description FROM %s u JOIN %s us ON u.id = us.note_id WHERE us.user_id = $1 AND us.note_id = $2", notesTable, usersNotesTable)
	err := p.db.Get(&note, query, userId, noteId)
	return note, err
}
func (p *NotePostgres) Delete(userId, noteId int) error {
	query := fmt.Sprintf("DELETE FROM %s u USING %s us WHERE  u.id = us.note_id AND us.user_id = $1 AND us.note_id = $2", notesTable, usersNotesTable)
	_, err := p.db.Exec(query, userId, noteId)
	return err
}
func (p *NotePostgres) Update(userId, noteId int, input notebook.UpdateNoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Date)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s u SET %s FROM %s us WHERE  u.id = us.note_id AND us.user_id = $%d AND us.note_id = $%d", notesTable, setQuery, usersNotesTable, argId, argId+1)
	args = append(args, userId, noteId)
	_, err := p.db.Exec(query, args...)
	return err
}
