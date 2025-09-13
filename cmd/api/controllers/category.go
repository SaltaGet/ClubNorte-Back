package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// CategoryGet godoc
//
// @Summary		CategoryGet
// @Description	CategoryGet obtener una categoria por ID
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			id	path		string	true	"Id de la categoria"
// @Success		200	{object}	schemas.Response{body=schemas.CategoryResponse}
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/api/v1/category/get/{id} [get]
func (c *CategoryController) CategoryGet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el id de la categoria", fmt.Errorf("se necesita el id de la categoria")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "Error al parsear el id de la categoria", err))
	}

	category, err := c.CategoryService.CategoryGetByID(uint(idUint))
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    category,
		Message: "Categoria obtenida con exito",
	})
}

// CategoryGetAll godoc
//
// @Summary		CategoryGet All
// @Description	CategoryGetAll obtener todas las categorias
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Success		200	{object}	schemas.Response{body=[]schemas.CategoryResponse}
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/api/v1/category/get_all [get]
func (c *CategoryController) CategoryGetAll(ctx *fiber.Ctx) error {
	categories, err := c.CategoryService.CategoryGetAll()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    categories,
		Message: "Categorias obtenidas con exito",
	})
}

// CategoryCreate godoc
//
// @Summary		CategoryCreate
// @Description	CategoryCreate crear una categoria
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			category_create	body		schemas.CategoryCreate	true	"Categoria a crear"
// @Success		200				{object}	schemas.Response{body=uint}
// @Failure		400				{object}	schemas.Response
// @Failure		401				{object}	schemas.Response
// @Failure		422				{object}	schemas.Response
// @Failure		404				{object}	schemas.Response
// @Failure		500				{object}	schemas.Response
// @Router			/api/v1/category/create [post]
func (c *CategoryController) CategoryCreate(ctx *fiber.Ctx) error {
	var categoryCreate *schemas.CategoryCreate
	if err := ctx.BodyParser(&categoryCreate); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := categoryCreate.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	id, err := c.CategoryService.CategoryCreate(categoryCreate)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(schemas.Response{
		Status:  true,
		Body:    id,
		Message: "Categoria creada exitosamente",
	})
}

// CategoryUpdate godoc
//
// @Summary		CategoryUpdate
// @Description	CategoryUpdate crear una categoria
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			category_update	body		schemas.CategoryUpdate	true	"Categoria a editar"
// @Success		200				{object}	schemas.Response
// @Failure		400				{object}	schemas.Response
// @Failure		401				{object}	schemas.Response
// @Failure		422				{object}	schemas.Response
// @Failure		404				{object}	schemas.Response
// @Failure		500				{object}	schemas.Response
// @Router			/api/v1/category/update [put]
func (c *CategoryController) CategoryUpdate(ctx *fiber.Ctx) error {
	var categoryUpdate *schemas.CategoryUpdate
	if err := ctx.BodyParser(&categoryUpdate); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := categoryUpdate.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	if err := c.CategoryService.CategoryUpdate(categoryUpdate); err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Categoria actualizada exitosamente",
	})
}

// CategoryDelete godoc
//
// @Summary		CategoryDelete
// @Description	CategoryDelete crear una categoria
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			id	path		string	true	"Categoria a eliminar por ID"
// @Success		200	{object}	schemas.Response
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/api/v1/category/delete/{id} [delete]
func (c *CategoryController) CategoryDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el id de la categoria", fmt.Errorf("se necesita el id de la categoria")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id ser un número", err))
	}

	if err := c.CategoryService.CategoryDelete(uint(idUint)); err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Categoria eliminada exitosamente",
	})
}
