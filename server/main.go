package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=finance_db sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}

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
	log.Fatal(app.Listen(":4000"))
}

//5:50 timecode stop
