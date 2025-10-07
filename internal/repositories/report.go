package repositories

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (r *MainRepository) ReportExcelGet() (any, error) {
	var products []*models.Product

	err := r.DB.Preload("Category").Find(&products).Error

	if err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener productos", err)
	}

	return products, err
}

// func (r *MainRepository) ReportMovementByDate(start, end time.Time) (*models.ReportMovement, error) {
// 	var income []*models.Income
// 	if err := r.DB.
// 		Where("created_at >= ? AND created_at < ?", start, end).
// 		Find(&income).Error; err != nil {
// 		return nil, schemas.ErrorResponse(500, "Error al obtener los ingresos", err)
// 	}

// 	var expense []*models.Expense
// 	if err := r.DB.
// 		Where("created_at >= ? AND created_at < ?", start, end).
// 		Find(&expense).Error; err != nil {
// 		return nil, schemas.ErrorResponse(500, "Error al obtener los egresos", err)
// 	}

// 	var incomeSportsCourts []*models.IncomeSportsCourts
// 	if err := r.DB.
// 		Where("created_at >= ? AND created_at < ?", start, end).
// 		Find(&incomeSportsCourts).Error; err != nil {
// 		return nil, schemas.ErrorResponse(500, "Error al obtener los ingresos por cancha", err)
// 	}

// 	var expenseBuy []*models.ExpenseBuy
// 	if err := r.DB.
// 		Where("created_at >= ? AND created_at < ?", start, end).
// 		Find(&expenseBuy).Error; err != nil {
// 		return nil, schemas.ErrorResponse(500, "Error al obtener los egresos por compra", err)
// 	}

// 	var PointSales []*models.PointSale
// 	if err := r.DB.Find(&PointSales).Error; err != nil {
// 		return nil, schemas.ErrorResponse(500, "Error al obtener los puntos de venta", err)
// 	}

// 	return &models.ReportMovement{
// 		Income:            income,
// 		Expense:           expense,
// 		IncomeSportsCourts: incomeSportsCourts,
// 		ExpenseBuy:        expenseBuy,
// 		PointSale:         PointSales,
// 	}, nil
// }

// func (r *MainRepository) ReportByPointSale(start, end time.Time) ([]ReportPointSaleFull, error) {
//     results := map[uint]*ReportPointSaleFull{}

//     // Ingresos
//     var incomeRows []struct {
//         PointSaleID uint
//         Total       float64
//     }
//     if err := r.DB.
//         Table("incomes").
//         Select("point_sale_id, SUM(total) as total").
//         Where("created_at >= ? AND created_at < ?", start, end).
//         Group("point_sale_id").
//         Scan(&incomeRows).Error; err != nil {
//         return nil, err
//     }

//     for _, row := range incomeRows {
//         if results[row.PointSaleID] == nil {
//             results[row.PointSaleID] = &ReportPointSaleFull{PointSaleID: row.PointSaleID}
//         }
//         results[row.PointSaleID].IncomeTotal = row.Total
//     }

//     // → repetir lo mismo para Expense, IncomeSportsCourts y ExpenseBuy
//     // y acumular en results[pointSaleID]

//     // Convertir a slice
//     var out []ReportPointSaleFull
//     for _, v := range results {
//         out = append(out, *v)
//     }
//     return out, nil
// }

func (r *MainRepository) ReportMovementByDate(start, end time.Time) (any, error) {
	var resultados []map[string]any
	// var resultados []*schemas.ResultadoPorDia

	query := `
	SELECT 
    ps.id as point_sale_id,
    ps.name as point_sale_name,
    DATE(mov.fecha) as fecha,
    COALESCE(SUM(CASE WHEN tipo = 'ingreso' THEN total ELSE 0 END),0) as total_ingresos,
    COALESCE(SUM(CASE WHEN tipo = 'egreso' THEN total ELSE 0 END),0) as total_egresos,
    COALESCE(SUM(CASE WHEN tipo = 'cancha' THEN total ELSE 0 END),0) as total_canchas,
    COALESCE(SUM(CASE WHEN tipo = 'compra' THEN total ELSE 0 END),0) as total_compras,
    COALESCE(SUM(CASE WHEN tipo IN ('ingreso','cancha') THEN total ELSE -total END),0) as balance
	FROM (
    SELECT created_at as fecha, total, 'ingreso' as tipo, point_sale_id
			FROM incomes
			WHERE created_at BETWEEN ? AND ?
			
			UNION ALL
			
			SELECT created_at as fecha, total, 'egreso' as tipo, point_sale_id
			FROM expenses
			WHERE created_at BETWEEN ? AND ?
			
			UNION ALL
			
			SELECT created_at as fecha, total, 'cancha' as tipo, point_sale_id
			FROM income_sports_courts
			WHERE created_at BETWEEN ? AND ?
	) AS mov
	JOIN point_sales ps ON ps.id = mov.point_sale_id
	WHERE mov.fecha BETWEEN ? AND ?
	GROUP BY ps.id, ps.name, DATE(mov.fecha)
	ORDER BY ps.id, fecha DESC

	`

	err := r.DB.Raw(query,
		start, end,
		start, end,
		start, end,
		start, end,
	).Scan(&resultados).Error

	grouped := make(map[string][]map[string]any)
	for _, row := range resultados {
		fecha := row["fecha"].(time.Time).Format("2006-01-02") // key simple
		grouped[fecha] = append(grouped[fecha], row)
	}

	// Transformar al formato esperado
	var result []map[string]any
	for fecha, movimientos := range grouped {
		result = append(result, map[string]any{
			"fecha":      fecha,
			"movimiento": movimientos,
		})
	}

	return result, err
}

func (r *MainRepository) ObtenerResumenPorDiaYPuntoVenta(fechaInicio, fechaFin time.Time) ([]schemas.ResultadoPorDiaYPuntoVenta, error) {
	var resultados []schemas.ResultadoPorDiaYPuntoVenta

	query := `
		SELECT 
			DATE(fecha) as fecha,
			point_sale_id,
			COALESCE(SUM(CASE WHEN tipo = 'ingreso' THEN total ELSE 0 END), 0) as total_ingresos,
			COALESCE(SUM(CASE WHEN tipo = 'egreso' THEN total ELSE 0 END), 0) as total_egresos,
			COALESCE(SUM(CASE WHEN tipo = 'cancha' THEN total ELSE 0 END), 0) as total_canchas,
			COALESCE(SUM(CASE WHEN tipo = 'compra' THEN total ELSE 0 END), 0) as total_compras,
			COALESCE(SUM(CASE WHEN tipo IN ('ingreso', 'cancha') THEN total ELSE -total END), 0) as balance
		FROM (
			SELECT created_at as fecha, total, point_sale_id, 'ingreso' as tipo
			FROM incomes
			WHERE created_at BETWEEN ? AND ?
			
			UNION ALL
			
			SELECT created_at as fecha, total, point_sale_id, 'egreso' as tipo
			FROM expenses
			WHERE created_at BETWEEN ? AND ?
			
			UNION ALL
			
			SELECT created_at as fecha, total, point_sale_id, 'cancha' as tipo
			FROM income_sports_courts
			WHERE created_at BETWEEN ? AND ?
			
			UNION ALL
			
			SELECT created_at as fecha, total, 0 as point_sale_id, 'compra' as tipo
			FROM expense_buys
			WHERE created_at BETWEEN ? AND ?
		) as movimientos
		GROUP BY DATE(fecha), point_sale_id
		ORDER BY fecha DESC, point_sale_id
	`

	err := r.DB.Raw(query,
		fechaInicio, fechaFin,
		fechaInicio, fechaFin,
		fechaInicio, fechaFin,
		fechaInicio, fechaFin,
	).Scan(&resultados).Error

	return resultados, err
}

// ========================
// CONSULTAS POR MES
// ========================

// ObtenerResumenPorMes obtiene todos los ingresos y egresos agrupados por mes para un punto de venta
// func ObtenerResumenPorMes(pointSaleID uint, anio int) ([]schemas.ResultadoPorMes, error) {
// 	var resultados []schemas.ResultadoPorMes

// 	query := `
// 		SELECT
// 			DATE_FORMAT(fecha, '%Y-%m') as mes,
// 			YEAR(fecha) as anio,
// 			COALESCE(SUM(CASE WHEN tipo = 'ingreso' THEN total ELSE 0 END), 0) as total_ingresos,
// 			COALESCE(SUM(CASE WHEN tipo = 'egreso' THEN total ELSE 0 END), 0) as total_egresos,
// 			COALESCE(SUM(CASE WHEN tipo = 'cancha' THEN total ELSE 0 END), 0) as total_canchas,
// 			COALESCE(SUM(CASE WHEN tipo = 'compra' THEN total ELSE 0 END), 0) as total_compras,
// 			COALESCE(SUM(CASE WHEN tipo IN ('ingreso', 'cancha') THEN total ELSE -total END), 0) as balance
// 		FROM (
// 			SELECT created_at as fecha, total, 'ingreso' as tipo
// 			FROM incomes
// 			WHERE point_sale_id = ? AND YEAR(created_at) = ?

// 			UNION ALL

// 			SELECT created_at as fecha, total, 'egreso' as tipo
// 			FROM expenses
// 			WHERE point_sale_id = ? AND YEAR(created_at) = ?

// 			UNION ALL

// 			SELECT created_at as fecha, total, 'cancha' as tipo
// 			FROM income_sports_courts
// 			WHERE point_sale_id = ? AND YEAR(created_at) = ?

// 			UNION ALL

// 			SELECT created_at as fecha, total, 'compra' as tipo
// 			FROM expense_buys
// 			WHERE YEAR(created_at) = ?
// 		) as movimientos
// 		GROUP BY DATE_FORMAT(fecha, '%Y-%m'), YEAR(fecha)
// 		ORDER BY mes DESC
// 	`

// 	err := db.Raw(query,
// 		pointSaleID, anio,
// 		pointSaleID, anio,
// 		pointSaleID, anio,
// 		anio,
// 	).Scan(&resultados).Error

// 	return resultados, err
// }
