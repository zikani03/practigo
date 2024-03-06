package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
}

var port int

func init() {
	flag.IntVar(&port, "port", 3000, "Port to run on")
}

func main() {
	flag.Parse()

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		p := Person{
			Name: "Joel",
		}

		return c.JSON(fiber.Map{
			"message": "Hello Tech Malawi",
			"person":  p,
		})
	})

	err := app.Listen(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}
}
