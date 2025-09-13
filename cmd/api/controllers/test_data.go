package controllers

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/database"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// TestDataCreate godoc
//
//	@Summary		TestDataCreate
//	@Description	TestDataCreate
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/test_data/create [post]
func TestDataCreate(c *fiber.Ctx) error {
	err := database.CreateTestData()
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Datos de prueba creados",
	})
}

// TestDataDelete godoc
//
//	@Summary		TestDataDelete
//	@Description	TestDataDelete
//	@Tags			Test
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/test_data/delete [delete]
func TestDataDelete(c *fiber.Ctx) error {
	err := database.DeleteTestData()
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Datos de prueba eliminados",
	})
}