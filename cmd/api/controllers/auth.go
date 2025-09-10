package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
//
// @Summary		Login user
// @Description	Login user required email and password
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			credentials	body		schemas.Login	true	"Credentials"
// @Success		200			{object}	schemas.Response
// @Failure		400			{object}	schemas.Response
// @Failure		401			{object}	schemas.Response
// @Failure		422			{object}	schemas.Response
// @Failure		404			{object}	schemas.Response
// @Failure		500			{object}	schemas.Response
// @Router			/v1/auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var credentials schemas.Login
	err := ctx.BodyParser(&credentials)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	err = credentials.Validate()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	token, err := c.AuthService.Login(&credentials)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}

	ctx.Cookie(cookie)

	// También podemos devolver un mensaje opcional en el body
	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Login exitoso, token enviado en cookie",
	})
}

// Logout godoc
//
// @Summary		Logout user
// @Description	Logout
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Success		200	{object}	schemas.Response
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/v1/auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Logout exitoso",
	})
}

// LoginPointSale godoc
//
// @Summary		LoginPointSale user
// @Description	LoginPointSale
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			point_sale_id	path		string	true	"point_sale_id"
// @Success		200				{object}	schemas.Response
// @Failure		400				{object}	schemas.Response
// @Failure		401				{object}	schemas.Response
// @Failure		422				{object}	schemas.Response
// @Failure		404				{object}	schemas.Response
// @Failure		500				{object}	schemas.Response
// @Router			/v1/auth/login_point_sale/{point_sale_id} [post]
func (c *AuthController) LoginPointSale(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemas.UserContext)

	var idParam = ctx.Params("point_sale_id")
	if idParam == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el id del punto de venta", fmt.Errorf("se necesita el id del punto de venta")))
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	token, err := c.AuthService.LoginPointSale(user.ID, uint(id))
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().AddDate(1, 0, 0),
	}

	ctx.Cookie(cookie)

	// También podemos devolver un mensaje opcional en el body
	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Login a Punto de venta exitoso, token enviado en cookie",
	})
}

// LogoutPointSale godoc
//
// @Summary		LogoutPointSale user
// @Description	LogoutPointSale
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Success		200	{object}	schemas.Response
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/v1/auth/logout_point_sale [post]
func (c *AuthController) LogoutPointSale(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemas.UserContext)

	token, err := c.AuthService.LogoutPointSale(user.ID)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	cookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().AddDate(1, 0, 0),
	}

	ctx.Cookie(cookie)

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Logout de Punto de venta exitoso, token enviado en cookie",
	})
}

// CurrentUser godoc
//
// @Summary		CurrentUser user
// @Description	CurrentUser
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Success		200	{object}	schemas.Response
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/v1/auth/current_user [get]
func (c *AuthController) CurrentUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemas.UserContext)

	userResponse, err := c.AuthService.CurrentUser(user.ID)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    userResponse,
		Message: "Usuario actual obtenido con exito",
	})
}

// CurrentPointSale godoc
//
// @Summary		CurrentPointSale point sale
// @Description	CurrentPointSale
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Success		200	{object}	schemas.Response
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/v1/auth/current_point_sale [get]
func (c *AuthController) CurrentPointSale(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)

	pointSaleResponse, err := c.AuthService.CurrentPointSale(pointSale.ID)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    pointSaleResponse,
		Message: "Punto de venta actual obtenido con exito",
	})
}
