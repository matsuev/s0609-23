package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/connect", connectHandler)

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
