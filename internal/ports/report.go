package ports

import (
	"time"

	// "github.com/DanielChachagua/Club-Norte-Back/internal/models"
	// "github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type ReportRepository interface {
	ReportExcelGet() (report any, err error)
	ReportMovementByDate(fromDate, toDate time.Time, form string) (report any, err error)
	ReportProfitableProducts() (report any, err error)
}

type ReportService interface {
	ReportExcelGet() (report any, err error)
	ReportMovementByDate(fromDate, toDate time.Time, form string) (report any, err error)
	ReportProfitableProducts() (report any, err error)
}