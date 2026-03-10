package main

import (
	"oat431/go-fiber-snippets-vol2/internal/bootstrap"
	"oat431/go-fiber-snippets-vol2/internal/config"
	"oat431/go-fiber-snippets-vol2/internal/routes"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	config.LoadEnvConfig()
	db := config.StartDatabase()
	defer db.Close()

	container := bootstrap.NewAppContainer()

	app := fiber.New()
	routes.SetupRoutes(app, container)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port :+ " + port + " is already in use")
	}
}
