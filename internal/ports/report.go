package ports

import (
	"time"

	// "github.com/DanielChachagua/Club-Norte-Back/internal/models"
	// "github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type ReportRepository interface {
	ReportExcelGet() (report any, err error)
	// ReportMovementByDate(fromDate, toDate time.Time) (report *models.ReportMovement, err error)
	ReportMovementByDate(fromDate, toDate time.Time) (report any, err error)
}

type ReportService interface {
	ReportExcelGet() (report any, err error)
	ReportMovementByDate(fromDate, toDate time.Time) (report any, err error)
	// ReportMovementByDate(fromDate, toDate time.Time) (report *schemas.ReportMovementResponse, err error)
}