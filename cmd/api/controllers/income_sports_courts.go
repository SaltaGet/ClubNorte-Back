package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// IncomeSportCourtGetByID godoc
//
//	@Summary		IncomeSportCourtGetByID
//	@Description	Obtener un ingreso por ID de una cancha
//	@Tags			IncomeSportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del ingreso"
//	@Success		200	{object}	schemas.Response{body=schemas.IncomeSportsCourtsResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/income_sport_court/get/{id} [get]
func (i *IncomeSportCourtController) IncomeSportCourtGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del ingreso", fmt.Errorf("se necesita el id del ingreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	income, err := i.IncomeSportCourtService.IncomeSportCourtGetByID(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    income,
		Message: "Ingreso obtenido exitosamente",
	})
}

// IncomeSportCourtGetByDate godoc
//
//	@Summary		IncomeSportCourtGetByDate
//	@Description	Obtener ingresos por fechas
//	@Tags			IncomeSportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			income_date	body		schemas.IncomeSportsCourtsDateRequest	true	"Fecha desde - hasta del ingreso"
//	@Param			page		query		int										false	"Número de página"				default(1)
//	@Param			limit		query		int										false	"Número de elementos por página"	default(10)
//	@Success		200			{object}	schemas.Response{body=[]schemas.IncomeSportsCourtsResponseDTO}
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/income_sport_court/get_by_date [post]
func (i *IncomeSportCourtController) IncomeSportCourtGetByDate(c *fiber.Ctx) error {
	var incomeDateRequest schemas.IncomeSportsCourtsDateRequest
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

	incomes, total, err := i.IncomeSportCourtService.IncomeSportCourtGetByDate(pointSale.ID, fromDate, toDate, page, limit)
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

// IncomeSportCourtCreate godoc
//
//	@Summary		IncomeSportCourtCreate
//	@Description	Crear un ingreso de una cancha
//	@Tags			IncomeSportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			income_create	body		schemas.IncomeSportsCourtsCreate	true	"Datos requeridos para crear un ingreso"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/income_sport_court/create [post]
func (i *IncomeSportCourtController) IncomeSportCourtCreate(c *fiber.Ctx) error {
	var incomeCreate schemas.IncomeSportsCourtsCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := incomeCreate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)
	user := c.Locals("user").(*schemas.UserContext)

	id, err := i.IncomeSportCourtService.IncomeSportCourtCreate(user.ID, pointSale.ID, &incomeCreate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    id,
		Message: "Ingreso creado exitosamente",
	})
}

// IncomeSportCourtUpdatePay godoc
//
//	@Summary		IncomeSportCourtUpdate
//	@Description	Editar un ingreso de una cancha
//	@Tags			IncomeSportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			income_sport_court_update_pay	body		schemas.IncomeSportsCourtsRestPay	true	"Datos requeridos para editar un ingreso"
//	@Success		200								{object}	schemas.Response
//	@Failure		400								{object}	schemas.Response
//	@Failure		401								{object}	schemas.Response
//	@Failure		422								{object}	schemas.Response
//	@Failure		404								{object}	schemas.Response
//	@Failure		500								{object}	schemas.Response
//	@Router			/api/v1/income_sport_court/update_pay [put]
func (i *IncomeSportCourtController) IncomeSportCourtUpdatePay(c *fiber.Ctx) error {
	var incomeSportCourtUpdatePay schemas.IncomeSportsCourtsRestPay
	if err := c.BodyParser(&incomeSportCourtUpdatePay); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}
	if err := incomeSportCourtUpdatePay.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)
	user := c.Locals("user").(*schemas.UserContext)

	err := i.IncomeSportCourtService.IncomeSportCourtUpdatePay(user.ID, pointSale.ID, &incomeSportCourtUpdatePay)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Ingreso actualizado exitosamente",
	})
}

// IncomeSportCourtDelete godoc
//
//	@Summary		IncomeSportCourtDelete
//	@Description	Eliminar un ingreso
//	@Tags			IncomeSportCourt
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
//	@Router			/api/v1/income_sport_court/delete/{id} [delete]
func (i *IncomeSportCourtController) IncomeSportCourtDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del ingreso", fmt.Errorf("se necesita el id del ingreso")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	err = i.IncomeSportCourtService.IncomeSportCourtDelete(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(200).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Ingreso eliminado exitosamente",
	})
}
