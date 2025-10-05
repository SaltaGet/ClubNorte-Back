package controllers

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
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

// ReportGetByDate godoc
//
//	@Summary		ReportGetByDate
//	@Description	Obtiene un reporte por fechas
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			dateRangeRequest	body		schemas.DateRangeRequest	true	"Rango de fechas"
//	@Success		200					{object}	schemas.Response{body=schemas.ReportMovementResponse}
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/api/v1/report/get_by_date [post]
func (c *ReportController) ReportMovementByDate(ctx *fiber.Ctx) error {
	var dateRangeRequest schemas.DateRangeRequest
	if err := ctx.BodyParser(&dateRangeRequest); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	fromDate, toDate, err := dateRangeRequest.GetParsedDates()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	report, err := c.ReportService.ReportMovementByDate(fromDate, toDate)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:    report,
		Message: "Reporte mensual obtenido con exito",
	})
}