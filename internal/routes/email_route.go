package routes

import (
	"oat431/go-fiber-snippets-vol2/internal/controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterEmailRoutes(router fiber.Router, controller controller.EmailController) {
	route := router.Group("/email")

	route.Get("/send", controller.SendEmail)
}
