package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func (c *InformController) InformGet(ctx *fiber.Ctx) error {
	inform, err := c.InformService.Inform()
	if err != nil {
		return err
	}

	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", "attachment; filename=inform.xlsx")
	ctx.Set("File-Name", "inform.xlsx")
	
	f, ok := inform.(*excelize.File)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error al crear el archivo")
	}

	if err := f.Write(ctx.Response().BodyWriter()); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return nil
}