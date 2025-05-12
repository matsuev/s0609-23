package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api", ApiHandler())

	if err := app.Listen(":8081"); err != nil {
		log.Fatalln(err)
	}
}

// ApiHandler ...
func ApiHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "test message",
		})
	}
}
