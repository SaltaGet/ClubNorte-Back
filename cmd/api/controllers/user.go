package controllers

import (
	"strconv"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

//	UserGetByID godoc
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
//	@Router			/v1/user/get/{id} [get]
func (u *UserController) UserGetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el id del usuario",
		})
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El id del usuario debe ser un número",
		})
	}

	user, err := u.UserService.UserGetByID(uint(idUint))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    user,
		Message: "Usuario obtenido con éxito",
	})	
}

//	UserGetByEmail godoc
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
//	@Router			/v1/user/get_by_email [get]
func (u *UserController) UserGetByEmail(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if email == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el email del usuario",
		})
	}

	user, err := u.UserService.UserGetByEmail(email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    user,
		Message: "Usuario obtenido con éxito",
	})	
}

//	UserGetAll godoc
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
//	@Router			/v1/user/get_all [get]
func (u *UserController) UserGetAll(ctx *fiber.Ctx) error {
	users, err := u.UserService.UserGetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    users,
		Message: "Usuarios obtenidos con éxito",
	})
}

//	UserCreate godoc
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
//	@Router			/v1/user/create [post]
func (u *UserController) UserCreate(ctx *fiber.Ctx) error {
	var userCreate *schemas.UserCreate
	if err := ctx.BodyParser(&userCreate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	} 

	if err := userCreate.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	userID, err := u.UserService.UserCreate(userCreate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    userID,
		Message: "Usuario creado con éxito",
	})
}

//	UserUpdate godoc
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
//	@Router			/v1/user/update [put]
func (u *UserController) UserUpdate(ctx *fiber.Ctx) error {
	var userUpdate *schemas.UserUpdate
	if err := ctx.BodyParser(&userUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	if err := userUpdate.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	if err := u.UserService.UserUpdate(userUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Usuario actualizado con éxito",
	})
}

//	UserDelete godoc
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
//	@Router			/v1/user/delete/{id} [delete]
func (u *UserController) UserDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "Se necesita el id del usuario",
		})
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: "El id del usuario debe ser un número",
		})
	}

	if err := u.UserService.UserDelete(uint(idUint)); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Usuario eliminado con éxito",
	})
}

func (u *UserController) UserUpdatePassword(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*schemas.UserContext)

	var updatePass *schemas.UserUpdatePassword
	if err := ctx.BodyParser(&updatePass); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	if err := updatePass.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	if err := u.UserService.UserUpdatePassword(user.ID, updatePass); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	// plantear logout luego de cambias pass

	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Contraseña actualizada con éxito",
	})
}

