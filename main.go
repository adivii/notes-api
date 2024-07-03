package main

import (
	"adivii/notes-api/configs"
	"adivii/notes-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create new echo object
	e := echo.New()
	db := configs.DatabaseInit()

	// Load Routes
	routes.NotesRoute(e, db)

	e.Logger.Fatal(e.Start(":9000"))
}
