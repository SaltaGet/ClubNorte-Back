package controllers

import (
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func (r *RegisterController) RegiterOpen(ctx *fiber.Ctx) error {
	var amountOpen schemas.RegisterOpen
	if err := ctx.BodyParser(&amountOpen); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := amountOpen.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	pointaSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	user := ctx.Locals("user").(*schemas.UserContext)

	err := r.RegisterService.RegisterOpen(pointaSale.ID, user.ID, amountOpen)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Apertura de caja realizada con exito",
	})
}

func (r *RegisterController) RegiterClose(ctx *fiber.Ctx) error {
	var amountClose schemas.RegisterClose
	if err := ctx.BodyParser(&amountClose); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := amountClose.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	pointaSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	user := ctx.Locals("user").(*schemas.UserContext)

	err := r.RegisterService.RegisterClose(pointaSale.ID, user.ID, amountClose)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Cierre de caja realizado con exito",
	})
}

func (r *RegisterController) RegiterInform(ctx *fiber.Ctx) error {
	dateInform := time.Now().UTC()

	pointaSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	user := ctx.Locals("user").(*schemas.UserContext)

	inform, err := r.RegisterService.RegisterInform(pointaSale.ID, user.ID, dateInform)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    inform,
		Message: fmt.Sprintf("Informe del %t obtenido con exito",dateInform),
	})
}