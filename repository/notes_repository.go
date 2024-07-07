package repository

import (
	"adivii/notes-api/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type NotesRepository struct {
	db *sqlx.DB
}

func NewNotesRepository(db *sqlx.DB) models.NotesRepository {
	return &NotesRepository{
		db: db,
	}
}

func (p *NotesRepository) CreateNotes(req models.Notes) (sql.Result, error) {
	query := fmt.Sprintf(`INSERT INTO "public"."notes" (title,"content",created_at,updated_at) VALUES ('%s', '%s', '%s', '%s')`, req.Title, req.Content, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	return p.db.Exec(query)
}

func (p *NotesRepository) GetAllNotes() ([]models.Notes, error) {
	var notes []models.Notes

	query := `SELECT id, title, "content", created_at, updated_at, labels_id FROM "public"."notes" WHERE is_deleted=false`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Notes

		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.LabelsId)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (p *NotesRepository) GetNotesById(id uuid.UUID) (*models.Notes, error) {
	var note models.Notes

	query := fmt.Sprintf(`SELECT id, title, "content", created_at, updated_at, labels_id FROM "public"."notes" WHERE id='%s' AND is_deleted=false`, id.String())
	row := p.db.QueryRow(query)

	err := row.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.LabelsId)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (p *NotesRepository) UpdateNotes(req models.Notes) (sql.Result, error) {
	query := fmt.Sprintf(`UPDATE "public"."notes" SET title='%s', "content"='%s', updated_at='%s' WHERE id='%s' AND is_deleted=false`, req.Title, req.Content, time.Now().Format(time.RFC3339), req.Id.String())
	return p.db.Exec(query)
}

func (p *NotesRepository) DeleteNotes(id uuid.UUID) (sql.Result, error) {
	query := fmt.Sprintf(`UPDATE "public"."notes" SET is_deleted=true WHERE id='%s' AND is_deleted=false`, id.String())
	return p.db.Exec(query)
}
