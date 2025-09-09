package controllers

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

//	RoleGetAll godoc
//
//	@Summary		RoleGetAll 
//	@Description	RoleGetAll obtener todos los roles
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response{body=[]schemas.RoleResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/v1/role/get_all [get]
func (r *RoleController) RoleGetAll(c *fiber.Ctx) error {
	roles, err := r.RoleService.RoleGetAll()
	if err != nil {
			return schemas.HandleError(c, err)
		}

	return c.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    roles,
		Message: "Roles obtenidos correctamente",
	}) 
}