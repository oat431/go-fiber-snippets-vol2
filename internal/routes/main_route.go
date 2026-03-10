package routes

import (
	"oat431/go-fiber-snippets-vol2/internal/bootstrap"
	"oat431/go-fiber-snippets-vol2/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	log.Info("Initializing routes...")
}

func SetupRoutes(app *fiber.App, container *bootstrap.AppContainer) {
	app.Use(middleware.GlobalLogger)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	RegisterHealthRoutes(v1)
	RegisterEmailRoutes(v1, *container.EmailController)
}
