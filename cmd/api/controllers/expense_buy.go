package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// ExpenseBuyGetByID godoc
//
//	@Summary		ExpenseBuyGetByID
//	@Description	Obtener un egreso de compra por ID
//	@Tags			ExpenseBuy
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del ingreso"
//	@Success		200	{object}	schemas.Response{body=schemas.ExpenseBuyResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/expense_buy/get/{id} [get]
func (i *ExpenseBuyController) ExpenseBuyGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del egreso", fmt.Errorf("se necesita el id del egreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	egreso, err := i.ExpenseBuyService.ExpenseBuyGetByID(uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    egreso,
		Message: "Egreso de compra obtenido exitosamente",
	})
}

// ExpenseBuyGetByDate godoc
//
//	@Summary		ExpenseBuyGetByDate
//	@Description	Obtener egresos de compra por fechas
//	@Tags			ExpenseBuy
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			expense_date	body		schemas.DateRangeRequest	true	"Fecha desde - hasta del egreso"
//	@Param			page			query		int							false	"Número de página"				default(1)
//	@Param			limit			query		int							false	"Número de elementos por página"	default(10)
//	@Success		200				{object}	schemas.Response{body=[]schemas.ExpenseBuyResponseDTO}
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/expense_buy/get_by_date [post]
func (i *ExpenseBuyController) ExpenseBuyGetByDate(c *fiber.Ctx) error {
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

	expenses, total, err := i.ExpenseBuyService.ExpenseBuyGetByDate(fromDate, toDate, page, limit)
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

// ExpenseBuyCreate godoc
//
//	@Summary		ExpenseBuyCreate
//	@Description	Crear un egreso de compra
//	@Tags			ExpenseBuy
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			expense_create	body		schemas.ExpenseBuyCreate	true	"Datos requeridos para crear un egreso de compra"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/expense_buy/create [post]
func (i *ExpenseBuyController) ExpenseBuyCreate(c *fiber.Ctx) error {
	var incomeCreate schemas.ExpenseBuyCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := incomeCreate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	user := c.Locals("user").(*schemas.UserContext)

	id, err := i.ExpenseBuyService.ExpenseBuyCreate(user.ID, &incomeCreate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    id,
		Message: "Egreso creado exitosamente",
	})
}

// ExpenseBuyDelte godoc
//
//	@Summary		ExpenseBuyDelte
//	@Description	Eliminar un egreso de compra
//	@Tags			ExpenseBuy
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del egreso de compra"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/expense_buy/delete/{id} [delete]
func (i *ExpenseBuyController) ExpenseBuyDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del egreso de compra", fmt.Errorf("se necesita el id del egreso de compra")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	err = i.ExpenseBuyService.ExpenseBuyDelete(uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Egreso de compra eliminado exitosamente",
	})
}
