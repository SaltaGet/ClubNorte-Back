package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// MovementStockGet godoc
//
//	@Summary		MovementStockGet
//	@Description	MovementStockGet Obtener un movimiento de stock por ID
//	@Tags			MovementStock
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"ID del movimiento de stock"
//	@Success		200	{object}	schemas.Response{body=schemas.MovementStockResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/movement_stock/get/{id} [get]
func (m *MovementStockController) MovementStockGet(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Se necesita el id del movimiento de stock", fmt.Errorf("se necesita el id del movimiento de stock")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	movementStock, err := m.MovementStockService.MovementStockGetByID(uint(idUint))
	if err != nil {
		return schemas.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    movementStock,
		Message: "Movimiento de stock obtenido correctamente",
	})
}

// MovementStockGetAll godoc
//
//	@Summary		MovementStockGetAll
//	@Description	MovementStockGetAll Obtener movimeintos de sotck por paginacion
//	@Tags			MovementStock
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			page	query		int	false	"Número de página"				default(1)
//	@Param			limit	query		int	false	"Número de elementos por página"	default(10)
//	@Success		200		{object}	schemas.Response{body=[]schemas.MovementStockResponse}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/api/v1/movement_stock/get_all [get]
func (m *MovementStockController) MovementStockGetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	movementsStock, total, err := m.MovementStockService.MovementStockGetAll(page, limit)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    map[string]interface{}{"movements": movementsStock, "total": total, "page": page, "limit": limit, "total_pages": totalPages},
		Message: "Movimiento de stock obtenido correctamente",
	})
}

// MovementStock godoc
//
//	@Summary		MovementStock
//	@Description	MovementStock movimiento de stock entre doposito y puntos de ventas
//	@Tags			MovementStock
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			movement_stock	body		schemas.MovementStock	true	"movimiento de stock"
//	@Success		200				{object}	schemas.Response
//	@Failure		400				{object}	schemas.Response
//	@Failure		401				{object}	schemas.Response
//	@Failure		422				{object}	schemas.Response
//	@Failure		404				{object}	schemas.Response
//	@Failure		500				{object}	schemas.Response
//	@Router			/api/v1/movement_stock/move [post]
func (m *MovementStockController) MoveStock(c *fiber.Ctx) error {
	user := c.Locals("user").(*schemas.UserContext)

	var movementStock schemas.MovementStock
	if err := c.BodyParser(&movementStock); err != nil {
		return schemas.HandleError(c, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := movementStock.Validate(); err != nil {
		return schemas.HandleError(c, err)
	}

	err := m.MovementStockService.MoveStock(user.ID, &movementStock)
	if err != nil {
		return schemas.HandleError(c, err)
	}

	if movementStock.FromType == "deposit" && movementStock.ToType == "point_sale" {
		select {
		case m.NotificationController.NotifyCh <- struct{}{}:
		default:
			fmt.Println("Canal de notificaciones lleno, no se pudo enviar notificación")
		}
	}


	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Movimiento de stock realizado correctamente",
	})
}
