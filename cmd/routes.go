package main

import (
	"fiber-started/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutess(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)
}
