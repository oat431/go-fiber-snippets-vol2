package bootstrap

import "github.com/gofiber/fiber/v3/log"

type AppContainer struct {
}

func NewAppContainer() *AppContainer {
	log.Info("Starting Bootstrap Container...")
	return &AppContainer{}
}
