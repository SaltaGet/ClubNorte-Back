package controllers

import (
	"fmt"
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// UserGetByID godoc
//
//	@Summary		UserGetByID
//	@Description	UserGetByID obtener usuario por ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"Id del usuario"
//	@Success		200	{object}	schemas.Response{body=schemas.UserResponse}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/user/get/{id} [get]
func (u *UserController) UserGetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "se necesita el id del usuario", fmt.Errorf("se necesita el id del usuario")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	user, err := u.UserService.UserGetByID(uint(idUint))
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    user,
		Message: "Usuario obtenido con éxito",
	})
}

// UserGetByEmail godoc
//
//	@Summary		UserGetByEmail
//	@Description	UserGetByEmail obtener usuario por Email
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			email	query		string	true	"email del usuario"
//	@Success		200		{object}	schemas.Response{body=schemas.UserResponse}
//	@Failure		400		{object}	schemas.Response
//	@Failure		401		{object}	schemas.Response
//	@Failure		422		{object}	schemas.Response
//	@Failure		404		{object}	schemas.Response
//	@Failure		500		{object}	schemas.Response
//	@Router			/api/v1/user/get_by_email [get]
func (u *UserController) UserGetByEmail(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "se necesita el email del usuario", fmt.Errorf("se necesita el email del usuario")))
	}

	user, err := u.UserService.UserGetByEmail(email)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    user,
		Message: "Usuario obtenido con éxito",
	})
}

// UserGetAll godoc
//
//	@Summary		UserGetAll
//	@Description	UserGetAll obtener todos los usuarios
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Success		200	{object}	schemas.Response{body=[]schemas.UserResponseDTO}
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/user/get_all [get]
func (u *UserController) UserGetAll(ctx *fiber.Ctx) error {
	users, err := u.UserService.UserGetAll()
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    users,
		Message: "Usuarios obtenidos con éxito",
	})
}

// UserCreate godoc
//
//	@Summary		UserCreate
//	@Description	UserCreate Crear un nuevo usuario
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			user_create	body		schemas.UserCreate	true	"user create"
//	@Success		200			{object}	schemas.Response
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/user/create [post]
func (u *UserController) UserCreate(ctx *fiber.Ctx) error {
	var userCreate *schemas.UserCreate
	if err := ctx.BodyParser(&userCreate); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := userCreate.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	userID, err := u.UserService.UserCreate(userCreate)
	if err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    userID,
		Message: "Usuario creado con éxito",
	})
}

// UserUpdate godoc
//
//	@Summary		UserUpdate
//	@Description	UserUpdate Editar un usuario
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			user_update	body		schemas.UserUpdate	true	"user create"
//	@Success		200			{object}	schemas.Response
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/user/update [put]
func (u *UserController) UserUpdate(ctx *fiber.Ctx) error {
	var userUpdate *schemas.UserUpdate
	if err := ctx.BodyParser(&userUpdate); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := userUpdate.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	if err := u.UserService.UserUpdate(userUpdate); err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Usuario actualizado con éxito",
	})
}

// UserDelete godoc
//
//	@Summary		UserDelete
//	@Description	UserDelete Eliminar un usuario por ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"Id del usuario"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/user/delete/{id} [delete]
func (u *UserController) UserDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "se necesita el id del usuario", fmt.Errorf("se necesita el id del usuario")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	if err := u.UserService.UserDelete(uint(idUint)); err != nil {
		return schemas.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Usuario eliminado con éxito",
	})
}

// UserUpdatePassword godoc
//
//	@Summary		UserUpdatePassword
//	@Description	UserUpdatePassword Actualizar la contraseña
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			pass_update	body		schemas.UserUpdatePassword	true	"update password"
//	@Success		200			{object}	schemas.Response
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/api/v1/user/update_password [put]
func (u *UserController) UserUpdatePassword(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemas.UserContext)

	var updatePass *schemas.UserUpdatePassword
	if err := ctx.BodyParser(&updatePass); err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "Error al parsear el cuerpo de la solicitud", err))
	}

	if err := updatePass.Validate(); err != nil {
		return schemas.HandleError(ctx, err)
	}

	if err := u.UserService.UserUpdatePassword(user.ID, updatePass); err != nil {
		return schemas.HandleError(ctx, err)
	}

	// plantear logout luego de cambias pass

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Contraseña actualizada con éxito",
	})
}

// UserUpdateIsActive godoc
//
//	@Summary		UserUpdateIsActive
//	@Description	Actualizar un usuario a activo o inactivo
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		CookieAuth
//	@Param			id	path		string	true	"id del usuario a activar o inactivar"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Failure		401	{object}	schemas.Response
//	@Failure		422	{object}	schemas.Response
//	@Failure		404	{object}	schemas.Response
//	@Failure		500	{object}	schemas.Response
//	@Router			/api/v1/user/update_is_active/{id} [put]
func (u *UserController) UserUpdateIsActive(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return schemas.HandleError(ctx, schemas.ErrorResponse(400, "se necesita el id del usuario", fmt.Errorf("se necesita el id del usuario")))
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return schemas.HandleError(ctx, schemas.ErrorResponse(422, "el id debe ser un número", err))
	}

	if err := u.UserService.UserUpdateIsActive(uint(idUint)); err != nil {
		return schemas.HandleError(ctx, err)
	}

	// plantear logout luego de cambias pass

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Contraseña actualizada con éxito",
	})
}
