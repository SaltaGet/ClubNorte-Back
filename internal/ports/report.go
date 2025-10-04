package ports

type ReportRepository interface {
	ReportExcelGet() (report any, err error)
	ReportMonthGet() (report any, err error)
}

type ReportService interface {
	ReportExcelGet() (report any, err error)
	ReportMonthGet() (report any, err error)
}