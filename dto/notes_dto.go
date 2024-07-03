package dto

import "github.com/labstack/echo/v4"

type NotesResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

type NotesRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
