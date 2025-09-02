package controllers

import (
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// DepositProductGetByID godoc
//
//	@Summary		DepositProductGetByID
//	@Description	DepositProductGetByID obtener un producto por ID del deposito
//	@Tags			Deposit
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"Id del producto"
//	@Success		200	{object}	schemas.Response{body=schemas.DepositResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/deposit/get/{id} [get]
func (d *DepositController) DepositGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el id del deposito",
		})
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El id del deposito no es valido",
		})
	}

	product, err := d.DepositService.DepositGetByID(uint(idUint))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El deposito no existe",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    product,
		Message: "Producto del deposito obtenido con exito",
	})
}

// DepositProductGetByCode godoc
//
//	@Summary		DepositProductGetByCode
//	@Description	DepositProductGetByCode obtener un producto por codigo del deposito
//	@Tags			Deposit
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			code	query		string	true	"codigo del producto"
//	@Success		200		{object}	schemas.Response{body=schemas.DepositResponse}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/v1/deposit/get_by_code [get]
func (d *DepositController) DepositGetByCode(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el codigo del producto",
		})
	}

	product, err := d.DepositService.DepositGetByCode(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    product,
		Message: "Producto obtenido correctamente",
	})
}

// DepositProductGetByName godoc
//
//	@Summary		DepositProductGetByName
//	@Description	DepositProductGetByName obtener productos por por similitud de nombre
//	@Tags			Deposit
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			name	query		string	true	"nombre por aproximacion del producto"
//	@Success		200		{object}	schemas.Response{body=schemas.DepositResponse}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/v1/deposit/get_by_name [get]
func (d *DepositController) DepositGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El nombre debe de tener al menos 3 caracteres",
		})
	}

	products, err := d.DepositService.DepositGetByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    products,
		Message: "Productos obtenidos correctamente",
	})
}

// DepositProductGetAll godoc
//
//	@Summary		DepositProductGetAll
//	@Description	DepositProductGetAll obtener productos por paginacion
//	@Tags			Deposit
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response{body=[]schemas.DepositResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/deposit/get_all [get]
func (d *DepositController) DepositGetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	products, total, err := d.DepositService.DepositGetAll(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]interface{}{"products": products, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Productos obtenidos correctamente",
	})
}

// DepositProductUpdateStock godoc
//
//	@Summary		DepositProductUpdateStock
//	@Description	DepositProductUpdateStock alctualizar stock de un producto del deposito
//	@Tags			Deposit
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			stock_update	body		schemas.DepositUpdateStock	true	"nombre por aproximacion del producto"
//	@Success		200				{object}	schemas.Response{body=[]schemas.DepositResponse}
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/v1/deposit/update_stock [put]
func (d *DepositController) DepositUpdateStock(c *fiber.Ctx) error {
	var stockUpdate schemas.DepositUpdateStock
	if err := c.BodyParser(&stockUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	if err := stockUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := d.DepositService.DepositUpdateStock(stockUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Producto actualizado correctamente",
	})
}
