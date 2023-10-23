package main

import (
	"log"

	"github.com/crimsoncoder42/gocrm-fiber/database"
	"github.com/crimsoncoder42/gocrm-fiber/lead"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead",lead.DeleteLead)

	// routes for notes 
	app.Get("/api/v1/lead/:id/notes", lead.GetNotesForLead)
	app.Post("/api/v1/lead/:id/note", lead.AddNoteToLead)
	app.Put("/api/v1/lead/:id/note/:noteId", lead.UpdateNote)
	app.Delete("/api/v1/lead/:id/note/:noteId", lead.DeleteNote)
}


func main() {

	app := fiber.New()
	initDatabase()
	defer database.Close()
	
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}

// initDatabase is a helper function to initialize the database and automigrate without import cycle.  
func initDatabase() {
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	database.DBConn.AutoMigrate(&lead.Lead{}, &lead.Note{})
}

