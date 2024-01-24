package main

import (
	"log"

	"github.com/Pongsakorn2341/fiber-gorm-rst/database"
	"github.com/Pongsakorn2341/fiber-gorm-rst/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	api := app.Group("/api/v1")

	routes.SetUpBaseRoutes(api)

	log.Fatal(app.Listen(":5050"))
}
