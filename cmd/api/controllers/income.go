package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// IncomeGetByID godoc
//
//	@Summary		IncomeGetByID
//	@Description	Obtener un ingreso por ID
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del ingreso"
//	@Success		200	{object}	schemas.Response{body=schemas.IncomeResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/income/get/{id} [get]
func (i *IncomeController) IncomeGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del ingreso", fmt.Errorf("se necesita el id del ingreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	income, err := i.IncomeService.IncomeGetByID(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    income,
		Message: "Ingreso obtenido exitosamente",
	})
}

// IncomeGetByDate godoc
//
//	@Summary		IncomeGetByDate
//	@Description	Obtener ingresos por fechas
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			income_date	body		schemas.DateRangeRequest	true	"Fecha desde - hasta del ingreso"
//	@Param			page		query		int							false	"Número de página"				default(1)
//	@Param			limit		query		int							false	"Número de elementos por página"	default(10)
//	@Success		200			{object}	schemas.Response{body=[]schemas.IncomeResponseDTO}
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/income/get_by_date [post]
func (i *IncomeController) IncomeGetByDate(c *fiber.Ctx) error {
	var incomeDateRequest schemas.DateRangeRequest
	if err := c.BodyParser(&incomeDateRequest); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	fromDate, toDate, err := incomeDateRequest.GetParsedDates()
	if err != nil {
		return schemas.HandleError(c, err)
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	incomes, total, err := i.IncomeService.IncomeGetByDate(pointSale.ID, fromDate, toDate, page, limit)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]any{"incomes": incomes, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Ingresos obtenidos exitosamente",
	})
}

// IncomeCreate godoc
//
//	@Summary		IncomeCreate
//	@Description	Crear un ingreso
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			income_create	body		schemas.IncomeCreate	true	"Datos requeridos para crear un ingreso"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/income/create [post]
func (i *IncomeController) IncomeCreate(c *fiber.Ctx) error {
	var incomeCreate schemas.IncomeCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := incomeCreate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)
	user := c.Locals("user").(*schemas.UserContext)

	id, err := i.IncomeService.IncomeCreate(user.ID, pointSale.ID, &incomeCreate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    id,
		Message: "Ingreso creado exitosamente",
	})
}

// IncomeDelte godoc
//
//	@Summary		IncomeDelte
//	@Description	Eliminar un ingreso
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del ingreso"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/income/delete/{id} [delete]
func (i *IncomeController) IncomeDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del ingreso", fmt.Errorf("se necesita el id del ingreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	err = i.IncomeService.IncomeDelete(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Ingreso eliminado exitosamente",
	})
}
