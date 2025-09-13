package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func (s *SportCourtController) SportCourtGetByID(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id de la cancha", fmt.Errorf("se necesita el id del punto de venta")))
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
		Message: "Punto de venta obtenido correctamente",
	})
}

func (s *SportCourtController) SportCourtGetByCode(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	code := c.Params("code")
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
		Message: "Punto de venta obtenido correctamente",
	})
}

func (s *SportCourtController) SportCourtGetAllByPointSale(c *fiber.Ctx) error {
	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	sportCourt, err := s.SportCourtService.SportCourtGetAllByPointSale(pointSale.ID)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Punto de venta obtenido correctamente",
	})
}

func (s *SportCourtController) SportCourtGetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return schemas.HandleError(c, err)
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	sportCourt, err := s.SportCourtService.SportCourtGetAll(page, limit)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Punto de venta obtenido correctamente",
	})
}

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

func (s *SportCourtController) SportCourtUpdate(c *fiber.Ctx) error {
	var sportUpdate schemas.SportCourtUpdate
	if err := c.BodyParser(&sportUpdate); err != nil {
		return schemas.HandleError(c, fmt.Errorf("error al parsear el body"))
	}
	if err := sportUpdate.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	pointSale := c.Locals("point_sale").(*schemas.PointSaleContext)

	err := s.SportCourtService.SportCourtCreate(pointSale.ID, &sportUpdate)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   nil,
		Message: "Cancha actualizada correctamente",
	})
}

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

	sportCourt, err := s.SportCourtService.SportCourtDelete(pointSale.ID, uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status: true,
		Body:   sportCourt,
		Message: "Punto de venta obtenido correctamente",
	})
}