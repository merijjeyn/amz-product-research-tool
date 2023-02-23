package main

import (
	"example/hello/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	db.InitializePostgreDB()
	db.ConnectToMongoDB()
	defer db.DisconnectFromMongoDB()

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
