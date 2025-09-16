package controllers

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// RegisterExistOpen godoc
//
//	@Summary		RegisterExistOpen
//	@Description	Verifica si existe apertura de caja
//	@Tags			Register
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/register/exist_open [get]
func (r *RegisterController) RegisterExistOpen(ctx *fiber.Ctx) error {
	pointaSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	
	existOpen, err := r.RegisterService.RegisterExistOpen(pointaSale.ID)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	var message string

	if existOpen {
		message = "Existe apertura de caja"
	} else {
		message = "No existe apertura de caja"
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    existOpen,
		Message: message,
	})
}

// RegisterOpen godoc
//
//	@Summary		RegisterOpen
//	@Description	Apertura de caja
//	@Tags			Register
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			amount_open	body		schemas.RegisterOpen	true	"Monto de apertura de caja"
//	@Success		200			{object}	schemas.Response
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/register/open [post]
func (r *RegisterController) RegisterOpen(ctx *fiber.Ctx) error {
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

// RegisterClose godoc
//
//	@Summary		RegisterClose
//	@Description	Cierre de caja
//	@Tags			Register
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			amount_close	body		schemas.RegisterClose	true	"Monto de cierre de caja"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/register/close [post]
func (r *RegisterController) RegisterClose(ctx *fiber.Ctx) error {
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

// RegisterInform godoc
//
//	@Summary		RegisterInform
//	@Description	Informes de caja
//	@Tags			Register
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			register_request	body		schemas.RegisterInformRequest	true	"Fechas de solicitud de informe"
//	@Success		200					{object}	schemas.Response
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/api/v1/register/inform [post]
func (r *RegisterController) RegiterInform(ctx *fiber.Ctx) error {
	var registerInformRequest schemas.RegisterInformRequest
	if err := ctx.BodyParser(&registerInformRequest); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	fromDate, toDate, err := registerInformRequest.GetParsedDates()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	pointaSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	user := ctx.Locals("user").(*schemas.UserContext)

	informs, err := r.RegisterService.RegisterInform(pointaSale.ID, user.ID, fromDate, toDate)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    informs,
		Message: "informes obtenidos con exito",
	})
}
