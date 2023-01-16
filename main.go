package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	env := godotenv.Load()
	if env != nil {
		fmt.Println(env)
	}

	port := ":" + os.Getenv("PORT")
	url := os.Getenv("URL")

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Redirect(url)
		return nil
	})

	err := app.Listen(port)
	if err != nil {
		fmt.Println(err)
	}
}
