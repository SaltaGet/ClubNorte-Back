package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// StockProductGet godoc
//
// @Summary		StockProductGetByID
// @Description	Obtener un producto del punto de venta por ID
// @Tags			PointSaleProduct
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			id	path		string	true	"Id del producto"
// @Success		200	{object}	schemas.Response{body=schemas.ProductResponse}
// @Failure		400	{object}	schemas.Response
// @Failure		401	{object}	schemas.Response
// @Failure		422	{object}	schemas.Response
// @Failure		404	{object}	schemas.Response
// @Failure		500	{object}	schemas.Response
// @Router			/api/v1/point_sale_product/get/{id} [get]
func (s *StockController) StockProductGetByID(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)
	
	productID := ctx.Params("id")
	if productID == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el id del producto", fmt.Errorf("se necesita el id del producto")))
	}

	productIDUint, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	product, err := s.StockService.StockProductGetByID(pointSale.ID, uint(productIDUint))
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    product,
		Message: "Producto obtenido correctamente",
	})
}

// StockProductGetByCode godoc
//
// @Summary		StockProductGetByCode
// @Description	Obtener un producto del punto de venta por Codigo
// @Tags			PointSaleProduct
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			code	query		string	true	"codigo del producto"
// @Success		200		{object}	schemas.Response{body=schemas.ProductResponse}
// @Failure		400		{object}	schemas.Response
// @Failure		401		{object}	schemas.Response
// @Failure		422		{object}	schemas.Response
// @Failure		404		{object}	schemas.Response
// @Failure		500		{object}	schemas.Response
// @Router			/api/v1/point_sale_product/get_by_code [get]
func (s *StockController) StockProductGetByCode(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)

	code := ctx.Query("code")
	if code == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el codigo del producto", fmt.Errorf("se necesita el codigo del producto")))
	}

	product, err := s.StockService.StockProductGetByCode(pointSale.ID, code)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    product,
		Message: "Producto obtenido correctamente",
	})
}

// StockProductGetByName godoc
//
// @Summary		StockProductGetByName
// @Description	Obtener un producto del punto de venta por nombre
// @Tags			PointSaleProduct
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			name	query		string	true	"nombre del producto"
// @Success		200		{object}	schemas.Response{body=schemas.ProductResponseDTO}
// @Failure		400		{object}	schemas.Response
// @Failure		401		{object}	schemas.Response
// @Failure		422		{object}	schemas.Response
// @Failure		404		{object}	schemas.Response
// @Failure		500		{object}	schemas.Response
// @Router			/api/v1/point_sale_product/get_by_name [get]
func (s *StockController) StockProductGetByName(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)

	name := ctx.Query("name")
	if len(name) < 3 {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el nombre del producto, minimo 3 caracteres", fmt.Errorf("se necesita el nombre del producto")))
	}

	products, err := s.StockService.StockProductGetByName(pointSale.ID, name)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    products,
		Message: "Productos obtenidos correctamente",
	})
}

// StockProductGetByCategory godoc
//
// @Summary		StockProductGetByCategory
// @Description	Obtener un producto del punto de venta por Id de categoria
// @Tags			PointSaleProduct
// @Accept			json
// @Produce		json
// @Security		CookieAuth
// @Param			category_id	path		string	true	"ID de la categoria"
// @Success		200			{object}	schemas.Response{body=schemas.ProductResponseDTO}
// @Failure		400			{object}	schemas.Response
// @Failure		401			{object}	schemas.Response
// @Failure		422			{object}	schemas.Response
// @Failure		404			{object}	schemas.Response
// @Failure		500			{object}	schemas.Response
// @Router			/api/v1/point_sale_product/get_by_category/{category_id} [get]
func (s *StockController) StockProductGetByCategoryID(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)

	categoryID := ctx.Params("category_id")
	if categoryID == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Se necesita el id de la categoria", fmt.Errorf("se necesita el id de la categoria")))
	}

	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	products, err := s.StockService.StockProductGetByCategoryID(pointSale.ID, uint(categoryIDUint))
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    products,
		Message: "Productos obtenidos correctamente",
	})
}

// StockProductGetAll godoc
//
//	@Summary		StockProductGetAll
//	@Description	Obtener todos los productos del punto de venta
//	@Tags			PointSaleProduct
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			page	query		int	false	"Número de página"				default(1)
//	@Param			limit	query		int	false	"Número de elementos por página"	default(10)
//	@Success		200		{object}	schemas.Response{body=[]schemas.ProductResponseDTO}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/api/v1/point_sale_product/get_all [get]
func (s *StockController) StockProductGetAll(ctx *fiber.Ctx) error {
	pointSale := ctx.Locals("point_sale").(*schemas.PointSaleContext)

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	products, total, err := s.StockService.StockProductGetAll(pointSale.ID, page, limit)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]any{"products": products, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Productos obtenidos correctamente",
	})
}
