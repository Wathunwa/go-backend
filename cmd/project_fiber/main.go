package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	mydb "github.com/Wathunwa/go-backend.git/db"
	project "github.com/Wathunwa/go-backend.git/service/project_fiber"
)

func main() {
	addr := ":22002"
	dsn := "postgres://postgres:fqu5gx7U7mEiMnhIlrlRIA1O4OoQFW2qhmEDuFtTq8rqfcgkcDMCCvor4T4uMsWf@157.245.202.214:27001/postgres"

	//connect DB
	pool, err := mydb.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer mydb.Close(pool)

	app := fiber.New(fiber.Config{})

	// health check
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok test trigger ": true, "service": "project_fiber"})
	})

	project.New(pool).Mount(app)

	log.Println("project fiber listening on ", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
