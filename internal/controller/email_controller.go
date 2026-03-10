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
	to := "oat431@outlook.com"
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
