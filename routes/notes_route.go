package routes

import (
	"adivii/notes-api/controllers"
	"adivii/notes-api/repository"
	"adivii/notes-api/usecases"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NotesRoute(e *echo.Echo, db *sqlx.DB) {
	notesRepository := repository.NewNotesRepository(db)
	notesUsecase := usecases.NewNotesUsecase(notesRepository)
	notesController := controllers.NewProductController(notesUsecase)

	e.GET("api/notes/", notesController.GetAllNotes)
	e.POST("api/notes/", notesController.CreateNotes)
}
