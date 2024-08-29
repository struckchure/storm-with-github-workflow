package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func ReturnCurrentDate() string {
	return time.Now().Format("2006-01-02 3:4:5 PM")
}

func RootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"app":      "gofiber",
		"datetime": ReturnCurrentDate(),
	})
}

func main() {
	app := fiber.New()

	app.Get("/", RootHandler)

	app.Listen(":9090")
}
