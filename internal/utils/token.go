package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *schemas.UserResponseToken, pointSale *schemas.PointSaleResponse) (string, error) {
	if user == nil {
		return "", fmt.Errorf("el usuario no puede estar vacio")
	}

	claims := jwt.MapClaims{
		"id":         user.ID,
		"email":      user.Email,
		"is_admin":   user.IsAdmin,
		"role_id":    user.Role.ID,
		"role":       user.Role.Name,
	}

	if pointSale != nil {
		claims["point_sale_id"] = pointSale.ID
		claims["point_sale"] = pointSale.Name
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("error al generar el token: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	cleanToken := CleanToken(tokenString)
	token, err := jwt.Parse(cleanToken, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, schemas.ErrorResponse(401, "Token inválido", err)
	}

	return token.Claims, nil
}

func CleanToken(bearerToken string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(bearerToken, prefix) {
		return strings.TrimPrefix(bearerToken, prefix)
	}
	return bearerToken
}
