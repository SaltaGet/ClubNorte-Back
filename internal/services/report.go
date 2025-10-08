package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"

	"github.com/xuri/excelize/v2"
)

func (s *ReportService) ReportExcelGet() (any, error) {
	inform, err := s.ReportRepository.ReportExcelGet()
	if err != nil {
		return nil, err
	}

	excel := excelize.NewFile()
	sheet := "Sheet1"

	excel.SetSheetName(sheet, "productos")
	sheet = "productos"

	// Encabezados
	headers := []string{"ID", "Code", "Name", "Description", "Price", "Category", "Notifier", "Min Amount"}
	for i, h := range headers {
		col := string(rune('A' + i))
		excel.SetCellValue(sheet, col+"1", h)
	}

	// Estilo encabezados
	headerStyle, _ := excel.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#4F81BD"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	excel.SetCellStyle(sheet, "A1", "H1", headerStyle)

	// Llenar filas con productos
	products := inform.([]*models.Product)
	categoryCount := make(map[string]int) // contador por categoría

	for i, product := range products {
		row := strconv.Itoa(i + 2)

		excel.SetCellValue(sheet, "A"+row, product.ID)
		excel.SetCellValue(sheet, "B"+row, product.Code)
		excel.SetCellValue(sheet, "C"+row, product.Name)
		if product.Description != nil {
			excel.SetCellValue(sheet, "D"+row, *product.Description)
		}
		excel.SetCellValue(sheet, "E"+row, product.Price)
		excel.SetCellValue(sheet, "F"+row, product.Category.Name)
		excel.SetCellValue(sheet, "G"+row, product.Notifier)
		excel.SetCellValue(sheet, "H"+row, product.MinAmount)

		// contar categorías
		categoryCount[product.Category.Name]++
	}

	// Ajustar ancho columnas
	for i := 0; i < len(headers); i++ {
		col := string(rune('A' + i))
		excel.SetColWidth(sheet, col, col, 20)
	}

	// Congelar primera fila
	excel.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	// -------------------------
	// Insertar tabla resumen por categoría
	// -------------------------
	resumeStart := len(products) + 3
	resumeRow := strconv.Itoa(resumeStart)

	excel.SetCellValue(sheet, "J"+resumeRow, "Category")
	excel.SetCellValue(sheet, "K"+resumeRow, "Count")

	r := resumeStart + 1
	for cat, count := range categoryCount {
		excel.SetCellValue(sheet, "J"+strconv.Itoa(r), cat)
		excel.SetCellValue(sheet, "K"+strconv.Itoa(r), count)
		r++
	}

	categoryRange := fmt.Sprintf("%s!$J$%d:$J$%d", sheet, resumeStart+1, r-1)
	countRange := fmt.Sprintf("%s!$K$%d:$K$%d", sheet, resumeStart+1, r-1)

	// Gráfico de torta simple compatible con excelize v2
	if err := excel.AddChart(sheet, "J2", &excelize.Chart{
		Type: excelize.Pie3D,
		Series: []excelize.ChartSeries{
			{
				Name:       "Productos por Categoría",
				Categories: categoryRange,
				Values:     countRange,
				// DataLabel:   excelize.ChartDataLabel{ShowVal: true},
				// ShowLabel:   true, // muestra los valores en cada porción
				// ShowPercent: true, // muestra porcentaje en cada porción
			},
		},
		Title: []excelize.RichTextRun{{Text: "Distribución de Productos por Categoría"}},
	}); err != nil {
		return nil, err
	}

	return excel, nil
}

// func (r *ReportService) ReportMovementByDate(fromDate, toDate time.Time) (*schemas.ReportMovementResponse, error) {
func (r *ReportService) ReportMovementByDate(fromDate, toDate time.Time, form string) (any, error) {
	report, err := r.ReportRepository.ReportMovementByDate(fromDate, toDate, form)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *ReportService) ReportProfitableProducts(start, end time.Time) ([]schemas.ReportProfitableProducts, error) {
	report, err := r.ReportRepository.ReportProfitableProducts(start, end)
	if err != nil {
		return nil, err
	}

	return report, nil
}
