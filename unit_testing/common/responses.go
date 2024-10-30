package common

import "github.com/gofiber/fiber/v2"

type HandlerResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HandlerErrorResponse(c *fiber.Ctx, code int, message string) error {
	r := &HandlerResponse{
		Message: message,
		Status:  code,
	}
	return c.Status(r.Status).JSON(r)
}