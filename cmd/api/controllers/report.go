package controllers

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"google.golang.org/genproto/googleapis/type/month"
)

func (c *ReportController) ReportExcelGet(ctx *fiber.Ctx) error {
	inform, err := c.ReportService.ReportExcelGet()
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

func (c *ReportController) ReportMonthGet(ctx *fiber.Ctx) error {
	var date string
	if date == "" {
		date = time.Now().Format("2006-01")
	}

	report, err := c.ReportService.ReportMonthGet()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:    report,
		Message: "Reporte mensual obtenido con exito",
	})
}