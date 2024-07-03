package controllers

import (
	"adivii/notes-api/dto"
	"adivii/notes-api/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotesController struct {
	notesUsecase models.NotesUsecase
}

func NewProductController(usecase models.NotesUsecase) NotesController {
	return NotesController{
		notesUsecase: usecase,
	}
}

func (p *NotesController) CreateNotes(c echo.Context) error {
	var notes models.Notes

	err := c.Bind(&notes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.NotesResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newNotes := dto.NotesRequest{
		Title:   notes.Title,
		Content: notes.Content,
	}

	result, err := p.notesUsecase.CreateNotes(newNotes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.NotesResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, dto.NotesResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func (p *NotesController) GetAllNotes(c echo.Context) error {
	result, err := p.notesUsecase.GetAllNotes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.NotesResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	fmt.Println(result)
	return c.JSON(http.StatusOK, dto.NotesResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}
