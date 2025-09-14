package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// SportCourtGetByID godoc
//
//	@Summary		SportCourtGetByID
//	@Description	Obtener una cancha por id
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID de la cancha"
//	@Success		200	{object}	schemas.Response{body=schemas.SportCourtResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/sport_court/get/{id} [get]
func (s *SportCourtController) SportCourtGetByID(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id de la cancha", fmt.Errorf("se necesita el id de la cancha")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	sportCourt, err := s.SportCourtService.SportCourtGetByID(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Cancha obtenida correctamente",
	})
}

// SportCourtGetByCode godoc
//
//	@Summary		SportCourtGetByCode
//	@Description	Obtener una cancha por id
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			code	query		string	true	"Codigo de la cancha"
//	@Success		200		{object}	schemas.Response{body=schemas.SportCourtResponse}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/api/v1/sport_court/get_by_code [get]
func (s *SportCourtController) SportCourtGetByCode(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	code := c.Query("code")
	if code == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "el codigo de la cancha es requerido", fmt.Errorf("el codigo de la cancha es requerido")))
	}

	sportCourt, err := s.SportCourtService.SportCourtGetByCode(pointSale.ID, code)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Cancha obtenida correctamente",
	})
}

// SportCourtGetAllByPointSale godoc
//
//	@Summary		SportCourtGetAllByPointSale
//	@Description	Obtener todas las canchas de un punto de venta
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response{body=[]schemas.SportCourtResponseDTO}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/sport_court/get_all_by_point_sale [get]
func (s *SportCourtController) SportCourtGetAllByPointSale(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	sportCourt, err := s.SportCourtService.SportCourtGetAllByPointSale(pointSale.ID)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Canchas obtenidas correctamente",
	})
}

// SportCourtGetAll godoc
//
//	@Summary		SportCourtGetAll
//	@Description	Obtener todas las canchas
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response{body=[]schemas.SportCourtResponseDTO}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/sport_court/get_all [get]
func (s *SportCourtController) SportCourtGetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return schemas.HandleError(c, err)
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	sportCourts, total, err := s.SportCourtService.SportCourtGetAll(page, limit)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]any{"sport_courts": sportCourts, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Canchas obtenidas correctamente",
	})
}

// SportCourtCreate godoc
//
//	@Summary		SportCourtCreate
//	@Description	Crear una cancha
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			sport_court_create	body		schemas.SportCourtCreate	true	"Parametros para crear una cancha"
//	@Success		200					{object}	schemas.Response{body=int}
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/api/v1/sport_court/create [post]
func (s *SportCourtController) SportCourtCreate(c *fiber.Ctx) error {
	var sportCreate schemas.SportCourtCreate
	if err := c.BodyParser(&sportCreate); err != nil {
		return schemas.HandleError(c, fmt.Errorf("error al parsear el body"))
	}
	if err := sportCreate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	id, err := s.SportCourtService.SportCourtCreate(pointSale.ID, &sportCreate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   id,
		Message: "Cancha creada correctamente",
	})
}

// SportCourtUpdate godoc
//
//	@Summary		SportCourtUpdate
//	@Description	Editar una cancha
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			sport_court_update	body		schemas.SportCourtUpdate	true	"Parametros para editar una cancha"
//	@Success		200					{object}	schemas.Response
//	@Failure		400					{object}	schemas.Response
//	@Failure		401					{object}	schemas.Response
//	@Failure		422					{object}	schemas.Response
//	@Failure		404					{object}	schemas.Response
//	@Failure		500					{object}	schemas.Response
//	@Router			/api/v1/sport_court/update [put]
func (s *SportCourtController) SportCourtUpdate(c *fiber.Ctx) error {
	var sportUpdate schemas.SportCourtUpdate
	if err := c.BodyParser(&sportUpdate); err != nil {
		return schemas.HandleError(c, fmt.Errorf("error al parsear el body"))
	}
	if err := sportUpdate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	err := s.SportCourtService.SportCourtUpdate(pointSale.ID, &sportUpdate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   nil,
		Message: "Cancha actualizada correctamente",
	})
}

// SportCourtDelete godoc
//
//	@Summary		SportCourtDelete
//	@Description	Eliminar una cancha por id
//	@Tags			SportCourt
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID de la cancha"
//	@Success		200	{object}	schemas.Response{body=schemas.SportCourtResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/sport_court/delete/{id} [post]
func (s *SportCourtController) SportCourtDelete(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id de la cancha", fmt.Errorf("se necesita el id del punto de venta")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	err = s.SportCourtService.SportCourtDelete(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   nil,
		Message: "Cancha eliminada correctamente",
	})
}