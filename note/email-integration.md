# Email Integration go fiber

1. Config the smtp server

```go
package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	SMTPHost       string
	SMTPPort       int
	SMTPUser       string
	SMTPPassword   string
	CodeExpiration time.Duration
}

func GetConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		port = 587
	}
	return &Config{
		SMTPHost:       os.Getenv("SMTP_HOST"),
		SMTPPort:       port,
		SMTPUser:       os.Getenv("SMTP_USER"),
		SMTPPassword:   os.Getenv("SMTP_PASS"),
		CodeExpiration: time.Minute * 1,
	}
}
```

2. Create a function to send email

```go
package service

import (
	"fmt"
	"net/smtp"
	"oat431/go-fiber-snippets-vol2/internal/config"
)

type SMTPService struct {
	config *config.Config
}

func NewSMTPService(config *config.Config) *SMTPService {
	return &SMTPService{config: config}
}

func (s *SMTPService) SendMail(to string) error {
	subject := "Subject: Go Email Integration \n"
	body := "This is a test email sent from a Go application using SMTP."
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth(
		"",
		s.config.SMTPUser,
		s.config.SMTPPassword,
		s.config.SMTPHost,
	)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.SMTPHost, s.config.SMTPPort),
		auth,
		s.config.SMTPUser,
		[]string{to},
		message,
	)
	if err != nil {
		return err
	}

	return nil
}

```

3. Use the SMTP service in your application

```go
package controller

import (
	"oat431/go-fiber-snippets-vol2/internal/service"
	"oat431/go-fiber-snippets-vol2/pkg/common"

	"github.com/gofiber/fiber/v3"
)

type EmailController struct {
	service service.SMTPService
}

func NewEmailController(service service.SMTPService) *EmailController {
	return &EmailController{service: service}
}

func (email *EmailController) SendEmail(c fiber.Ctx) error {
	to := "user@mailserver.domain"
	err := email.service.SendMail(to)
	if err != nil {
		res := common.ResponseDTO[string]{
			Data:   nil,
			Status: common.ERROR,
			Error: &common.ResponseDTOError{
				Message:   "Failed to send email: " + err.Error(),
				HttpCode:  fiber.StatusInternalServerError,
				ErrorCode: "EMAIL_SEND_ERROR",
			},
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	sentMail := "Email sent successfully"
	res := common.ResponseDTO[string]{
		Data:   &sentMail,
		Status: common.SUCCESS,
		Error:  nil,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
```

4. register api like a normal api

bootstrap/api_container.go
```go
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
```

routes/email_route.go
```go
package routes

import (
	"oat431/go-fiber-snippets-vol2/internal/controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterEmailRoutes(router fiber.Router, controller controller.EmailController) {
	route := router.Group("/email")

	route.Get("/send", controller.SendEmail)
}
```

routes/main_route.go
```go
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

	RegisterEmailRoutes(v1, *container.EmailController)
}
```

main.go
```go
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

	container := bootstrap.NewAppContainer()

	app := fiber.New()
	routes.SetupRoutes(app, container)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port :+ " + port + " is already in use")
	}
}
```
