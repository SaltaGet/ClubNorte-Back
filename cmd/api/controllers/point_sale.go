package controllers

import (
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/logging"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

//  PointSaleGet godoc
//	@Summary		PointSaleGet
//	@Description	PointSaleGet
//	@Tags			PointSale
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID Point Sale"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		403	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/point_sale/get/{id} [get]
func (p *PointSaleController) PointSaleGet(c *fiber.Ctx) error {
	var idParam = c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el id del punto de venta",
		})
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El ID proporcionado no es válido",
		})
	}

	pointSale, err := p.PointSaleService.PointSaleGet(uint(id))
	if err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al obtener el punto de venta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    pointSale,
		Message: "Punto de venta obtenido correctamente",
	})
}

//  PointSaleGetAll godoc
//	@Summary		PointSaleGetAll
//	@Description	PointSaleGetAll
//	@Tags			PointSale
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/point_sale/get_all [get]
func (p *PointSaleController) PointSaleGetAll(c *fiber.Ctx) error {
	pointsSales, err := p.PointSaleService.PointSaleGetAll()
	if err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al obtener los puntos de venta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    pointsSales,
		Message: "Puntos de venta obtenidos correctamente",
	})
}

//  PointSaleCreate godoc
//	@Summary		PointSaleCreate
//	@Description	PointSaleCreate
//	@Tags			PointSale
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			point_sale_create	body		schemas.PointSaleCreate	true	"PointSaleCreate"
//	@Success		200					{object}	schemas.Response
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/v1/point_sale/create [post]
func (p *PointSaleController) PointSaleCreate(c *fiber.Ctx) error {
	var pointSaleCreate schemas.PointSaleCreate
	if err := c.BodyParser(&pointSaleCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al parsear el cuerpo de la solicitud",
		})
	}

	if err := pointSaleCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	pointSale, err := p.PointSaleService.PointSaleCreate(&pointSaleCreate)
	if err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear el punto de venta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    pointSale,
		Message: "Punto de venta creado correctamente",
	})
}

//  PointSaleUpdate godoc
//	@Summary		PointSaleUpdate
//	@Description	PointSaleUpdate
//	@Tags			PointSale
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			point_sale_update	body		schemas.PointSaleUpdate	true	"PointSaleUpdate"
//	@Success		200					{object}	schemas.Response
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/v1/point_sale/update [put]
func (p *PointSaleController) PointSaleUpdate(c *fiber.Ctx) error {
	var pointSaleUpdate schemas.PointSaleUpdate
	if err := c.BodyParser(&pointSaleUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al parsear el cuerpo de la solicitud",
		})
	}

	if err := pointSaleUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := p.PointSaleService.PointSaleUpdate(&pointSaleUpdate)
	if err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al actualizar el punto de venta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Punto de venta actualizado correctamente",
	})
}

//  PointSaleDelete godoc
//	@Summary		PointSaleDelete
//	@Description	PointSaleDelete
//	@Tags			PointSale
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID Point Sale"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/point_sale/delete/{id} [delete]
func (p *PointSaleController) PointSaleDelete(c *fiber.Ctx) error {
	var idParam = c.Params("id")
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el id del punto de venta",
		})
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El ID proporcionado no es válido",
		})
	}

	err = p.PointSaleService.PointSaleDelete(uint(id))
	if err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al eliminar el punto de venta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Punto de venta eliminado correctamente",
	})
}
