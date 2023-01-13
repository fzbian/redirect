package main

import (
	"fmt"
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

	app.Get("/", func(c *fiber.Ctx) error {
		c.Redirect(url)
		return nil
	})

	err := app.Listen(port)
	if err != nil {
		fmt.Println(err)
	}
}
