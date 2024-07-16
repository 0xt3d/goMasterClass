package main

import(
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Type("html")
		return c.SendString("Hello, World")
	})
	app.Listen(":3030")
}