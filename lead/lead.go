package lead

import (
	"strconv"
	"github.com/crimsoncoder42/gocrm-fiber/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

)

// Lead represents a potential customer or client.

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
	Notes   []Note `gorm:"foreignKey:LeadID"`
}

// Note represents additional details or comments about a lead.
type Note struct {
	gorm.Model
	Content string `json:"content"`
	LeadID  uint   `json:"lead_id"`
}
// GetLeads retrieves all the leads from the database.
func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

// GetLead retrieves a single lead by its ID.

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

// NewLead creates a new lead in the database.

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&lead)
	return c.JSON(lead)
}

// DeleteLead removes a lead from the database by its ID.

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully Deleted")
}

// GetNotesForLead retrieves all notes associated with a given lead ID.

func GetNotesForLead(c *fiber.Ctx) error {
	leadID := c.Params("id")
	db := database.DBConn

	var notes []Note
	db.Where("lead_id = ?", leadID).Find(&notes)

	return c.JSON(notes)
}

// AddNoteToLead creates a new note and associates it with a lead.

func AddNoteToLead(c *fiber.Ctx) error {
	leadIDStr := c.Params("id")
	db := database.DBConn

	note := new(Note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// Convert leadID from string to uint
	leadID, err := strconv.Atoi(leadIDStr)
	if err != nil {
		return c.Status(400).SendString("Invalid Lead ID")
	}

	note.LeadID = uint(leadID)
	db.Create(&note)

	return c.JSON(note)
}

// UpdateNote modifies the content of a note associated with a lead.
func UpdateNote(c *fiber.Ctx) error {
	leadID := c.Params("id")
	noteID := c.Params("noteId")
	db := database.DBConn

	var note Note
	if err := db.Where("ID = ? AND lead_id = ?", noteID, leadID).First(&note).Error; err != nil {
		return c.Status(404).SendString("Note not found")
	}

	if err := c.BodyParser(&note); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Save(&note)

	return c.JSON(note)
}

// DeleteNote removes a note associated with a lead from the database.

func DeleteNote(c *fiber.Ctx) error {
	leadID := c.Params("id")
	noteID := c.Params("noteId")
	db := database.DBConn

	db.Where("ID = ? AND lead_id = ?", noteID, leadID).Delete(&Note{})

	return c.SendString("Note successfully Deleted")
}


