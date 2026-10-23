package main

import (
	"log"

	"github.com/glorse127/fin_tracker/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	cfg := config.Load()
	db, err := sqlx.Connect("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		var result int
		err := db.Get(&result, "SELECT 1")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"hello":    "world",
			"database": "available",
		})
	})
	log.Fatal(app.Listen(":" + cfg.Port))
}
