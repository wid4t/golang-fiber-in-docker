package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	api := app.Group("/module/partner", logger.New())

	api.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, i'm from golang-fiber-in-docker")
	})

	log.Fatal(app.Listen(":3000"))
}
