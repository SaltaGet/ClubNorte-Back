package schemas

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/logging"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool   `json:"status"`
	Body    any    `json:"body"`
	Message string `json:"message"`
}

func HandleError(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	// Si es error conocido
	if errResp, ok := err.(*ErrorStruc); ok {
		logging.ERROR("Error: %s", errResp.Err.Error())
		return ctx.Status(errResp.StatusCode).JSON(Response{
			Status:  false,
			Body:    nil,
			Message: errResp.Message,
		})
	}

	// Si es error genérico
	logging.ERROR("Error: %s", err.Error())
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  false,
		Body:    nil,
		Message: "Error interno",
	})
}
