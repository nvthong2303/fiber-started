package handlers

import (
	"fiber-started/database"
	"fiber-started/models"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "List of Facts",
		"Subtitle": "Facts for funminutes with friends!",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "Add a new Fact",
		"Subtitle": "Add a new fact to the list!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact Added Successfully",
		"Subtitle": "Your fact has been added to the list!",
	})
}
