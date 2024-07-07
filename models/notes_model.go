package models

import (
	"adivii/notes-api/dto"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Notes struct {
	Id        uuid.UUID     `json:"id"`
	Title     string        `json:"title" validate:"required"`
	Content   string        `json:"content" validate:"required"`
	CreatedAt time.Time     `json:"created_at" validate:"required"`
	UpdatedAt time.Time     `json:"updated_at" validate:"required"`
	LabelsId  uuid.NullUUID `json:"labels_id"`
}

type NotesRepository interface {
	CreateNotes(req Notes) (sql.Result, error)
	GetAllNotes() ([]Notes, error)
	GetNotesById(id uuid.UUID) (*Notes, error)
	UpdateNotes(req Notes) (sql.Result, error)
	DeleteNotes(id uuid.UUID) (sql.Result, error)
}

type NotesUsecase interface {
	CreateNotes(req dto.NotesRequest) (sql.Result, error)
	GetAllNotes() ([]Notes, error)
	GetNotesById(id uuid.UUID) (*Notes, error)
	UpdateNotes(req dto.NotesRequest, id uuid.UUID) (sql.Result, error)
	DeleteNotes(id uuid.UUID) (sql.Result, error)
}
