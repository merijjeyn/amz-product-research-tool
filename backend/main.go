package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Print("Cannot open database connection: " + err.Error())
		return
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/login", login)
	app.Listen(":8080")
}
