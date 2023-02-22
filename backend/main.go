package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	// ENDPOINTS
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", login)
	app.Post("/analyse_search_terms", analyseSearchTerms)

	app.Listen(":8080")
}
