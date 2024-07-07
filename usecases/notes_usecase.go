package usecases

import (
	"adivii/notes-api/dto"
	"adivii/notes-api/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
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

func (p NotesUsecase) GetNotesById(id uuid.UUID) (*models.Notes, error) {
	return p.notesRepository.GetNotesById(id)
}

func (p NotesUsecase) UpdateNotes(req dto.NotesRequest, id uuid.UUID) (sql.Result, error) {
	newNotes := models.Notes{
		Id:        id,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return p.notesRepository.UpdateNotes(newNotes)
}

func (p NotesUsecase) DeleteNotes(id uuid.UUID) (sql.Result, error) {
	return p.notesRepository.DeleteNotes(id)
}
