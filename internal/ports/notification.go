package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type NotificationRepository interface {
	NotificationStock() ([]*models.Product, error)
}

type NotificationService interface {
	NotificationStock() ([]*schemas.ProductSimpleResponse, error)
}
