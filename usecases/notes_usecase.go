package usecases

import (
	"adivii/notes-api/dto"
	"adivii/notes-api/models"
	"database/sql"
	"time"
)

type NotesUsecase struct {
	notesRepository models.NotesRepository
}

func NewNotesUsecase(repo models.NotesRepository) models.NotesUsecase {
	return NotesUsecase{
		notesRepository: repo,
	}
}

func (p NotesUsecase) CreateNotes(req dto.NotesRequest) (sql.Result, error) {
	newNotes := models.Notes{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return p.notesRepository.CreateNotes(newNotes)
}

func (p NotesUsecase) GetAllNotes() ([]models.Notes, error) {
	return p.notesRepository.GetAllNotes()
}
