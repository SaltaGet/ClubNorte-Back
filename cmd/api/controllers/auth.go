package controllers

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

//	 Login godoc
//	@Summary		Login user
//	@Description	Login user required email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		schemas.Login	true	"Credentials"
//	@Success		200			{object}	schemas.Response
//	@Failure		400			{object}	schemas.Response
//	@Failure		401			{object}	schemas.Response
//	@Failure		422			{object}	schemas.Response
//	@Failure		404			{object}	schemas.Response
//	@Failure		500			{object}	schemas.Response
//	@Router			/v1/auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var credentials schemas.Login
	err := ctx.BodyParser(&credentials)
	if err != nil {
		return err
	}

	err = credentials.Validate()
	if err != nil {
		return err
	}

	token, err := c.AuthService.Login(&credentials)
	if err != nil {
		return err
	}

	cookie := &fiber.Cookie{
		Name:     "access_token",                 // Nombre de la cookie
		Value:    token,                          // El JWT
		HTTPOnly: true,                           // No accesible desde JS → más seguro
		Secure:   true,                           // Solo HTTPS
		SameSite: "Strict",                       // Evita envío cross-site
	}

	ctx.Cookie(cookie) // Enviamos la cookie al cliente

	// También podemos devolver un mensaje opcional en el body
	return ctx.Status(fiber.StatusOK).JSON(schemas.Response{
		Status:  true,
		Body:    nil,
		Message: "Login exitoso, token enviado en cookie",
	})
}
