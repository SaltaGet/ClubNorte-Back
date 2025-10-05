package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// ExpenseGetByID godoc
//
//	@Summary		ExpenseGetByID
//	@Description	Obtener un egreso por ID
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del ingreso"
//	@Success		200	{object}	schemas.Response{body=schemas.ExpenseResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/expense/get/{id} [get]
func (i *ExpenseController) ExpenseGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del egreso", fmt.Errorf("se necesita el id del egreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	egreso, err := i.ExpenseService.ExpenseGetByID(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    egreso,
		Message: "Egreso obtenido exitosamente",
	})
}

// ExpenseGetByDate godoc
//
//	@Summary		ExpenseGetByDate
//	@Description	Obtener egresos por fechas
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			expense_date	body		schemas.DateRangeRequest	true	"Fecha desde - hasta del egreso"
//	@Param			page			query		int							false	"Número de página"				default(1)
//	@Param			limit			query		int							false	"Número de elementos por página"	default(10)
//	@Success		200				{object}	schemas.Response{body=[]schemas.ExpenseResponseDTO}
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/expense/get_by_date [post]
func (i *ExpenseController) ExpenseGetByDate(c *fiber.Ctx) error {
	var expenseDateRequest schemas.DateRangeRequest
	if err := c.BodyParser(&expenseDateRequest); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	fromDate, toDate, err := expenseDateRequest.GetParsedDates()
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

	expenses, total, err := i.ExpenseService.ExpenseGetByDate(pointSale.ID, fromDate, toDate, page, limit)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]any{"expenses": expenses, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Egresos obtenidos exitosamente",
	})
}

// ExpenseCreate godoc
//
//	@Summary		ExpenseCreate
//	@Description	Crear un egreso
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			expense_create	body		schemas.ExpenseCreate	true	"Datos requeridos para crear un egreso"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/expense/create [post]
func (i *ExpenseController) ExpenseCreate(c *fiber.Ctx) error {
	var incomeCreate schemas.ExpenseCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := incomeCreate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)
	user := c.Locals("user").(*schemas.UserContext)

	id, err := i.ExpenseService.ExpenseCreate(user.ID, pointSale.ID, &incomeCreate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    id,
		Message: "Egreso creado exitosamente",
	})
}

// ExpenseDelte godoc
//
//	@Summary		ExpenseDelte
//	@Description	Eliminar un egreso
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del egreso"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/expense/delete/{id} [delete]
func (i *ExpenseController) ExpenseDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del egreso", fmt.Errorf("se necesita el id del egreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	err = i.ExpenseService.ExpenseDelete(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Egreso eliminado exitosamente",
	})
}
