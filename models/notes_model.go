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
	DeletedAt time.Time     `json:"deleted_at"`
	LabelsId  uuid.NullUUID `json:"labels_id"`
}

type NotesRepository interface {
	CreateNotes(req Notes) (sql.Result, error)
	GetAllNotes() ([]Notes, error)
}

type NotesUsecase interface {
	CreateNotes(req dto.NotesRequest) (sql.Result, error)
	GetAllNotes() ([]Notes, error)
}
