package bootstrap

import (
	"oat431/go-fiber-snippets-vol2/internal/config"
	"oat431/go-fiber-snippets-vol2/internal/controller"
	"oat431/go-fiber-snippets-vol2/internal/service"

	"github.com/gofiber/fiber/v3/log"
)

type AppContainer struct {
	EmailController *controller.EmailController
}

func NewAppContainer() *AppContainer {
	emailService := service.NewSMTPService(config.GetConfig())
	emailController := controller.NewEmailController(*emailService)

	log.Info("Starting Bootstrap Container...")
	return &AppContainer{
		EmailController: emailController,
	}
}
