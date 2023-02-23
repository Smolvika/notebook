package service

import (
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(userId int, note notebook.Note) (int, error) {
	return s.repo.Create(userId, note)
}

func (s *NoteService) GetAll(userId int) ([]notebook.Note, error) {
	return s.repo.GetAll(userId)
}
func (s *NoteService) GetById(userId, noteId int) (notebook.Note, error) {
	return s.repo.GetById(userId, noteId)
}

func (s *NoteService) Delete(userId, noteId int) error {
	return s.repo.Delete(userId, noteId)
}
func (s *NoteService) Update(userId, noteId int, input notebook.UpdateNoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, noteId, input)
}
