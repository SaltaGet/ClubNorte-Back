package database

import (
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
)

func CreateTestData() error {
	// 2. Crear Categorías
	categories := []models.Category{
		{ID: 1, Name: "Bebidas"},
		{ID: 2, Name: "Snacks"},
		{ID: 3, Name: "Equipamiento Deportivo"},
		{ID: 4, Name: "Accesorios"},
		{ID: 5, Name: "Comida"},
	}

	for _, category := range categories {
		dbTest.FirstOrCreate(&category, models.Category{ID: category.ID})
	}

	// 3. Crear Puntos de Venta
	pointSales := []models.PointSale{
		{
			ID:          1,
			Name:        "Punto Central",
			Description: stringPtr("Punto de venta principal del complejo"),
		},
		{
			ID:          2,
			Name:        "Cancha Norte",
			Description: stringPtr("Kiosco ubicado en las canchas del sector norte"),
		},
		{
			ID:          3,
			Name:        "Área de Descanso",
			Description: stringPtr("Punto de venta en la zona de descanso"),
		},
	}

	for _, pointSale := range pointSales {
		dbTest.FirstOrCreate(&pointSale, models.PointSale{ID: pointSale.ID})
	}

	// 4. Crear Canchas Deportivas
	sportsCourts := []models.SportsCourt{
		{
			ID:          1,
			Code:        "CANCHA-01",
			Name:        "Cancha Fútbol A",
			Description: stringPtr("Cancha de fútbol principal con césped sintético"),
		},
		{
			ID:          2,
			Code:        "CANCHA-02",
			Name:        "Cancha Fútbol B",
			Description: stringPtr("Cancha de fútbol secundaria"),
		},
		{
			ID:          3,
			Code:        "CANCHA-03",
			Name:        "Cancha Tenis 1",
			Description: stringPtr("Cancha de tenis con superficie de polvo de ladrillo"),
		},
		{
			ID:          4,
			Code:        "CANCHA-04",
			Name:        "Cancha Básquet",
			Description: stringPtr("Cancha de básquetbol techada"),
		},
	}

	for _, court := range sportsCourts {
		dbTest.FirstOrCreate(&court, models.SportsCourt{ID: court.ID})
	}

	var sportsCourt1, sportsCourt2, sportsCourt3, sportsCourt4 models.SportsCourt
	var pointSport1, pointSport2, pointSport3 models.PointSale

	dbTest.First(&sportsCourt1, 1)
	dbTest.First(&sportsCourt2, 2)
	dbTest.First(&sportsCourt3, 3)
	dbTest.First(&sportsCourt4, 4)
	dbTest.First(&pointSport1, 1)
	dbTest.First(&pointSport2, 2)
	dbTest.First(&pointSport3, 3)

	// Asignar usuarios a puntos de venta
	dbTest.Model(&pointSport1).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt1})
	dbTest.Model(&pointSport2).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt2})
	dbTest.Model(&pointSport3).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt3, sportsCourt4})

	// 5. Crear Usuarios
	users := []models.User{
		{
			ID:        2,
			FirstName: "Juan Carlos",
			LastName:  "Pérez",
			Address:   stringPtr("Av. San Martín 123"),
			Cellphone: stringPtr("11-2345-6789"),
			Email:     "1@mail.com",
			Username:  "jperez",
			Password:  "1", // En producción usar hash real
			IsActive:  true,
			IsAdmin:   false,
			RoleID:    1,
		},
		{
			ID:        3,
			FirstName: "María Elena",
			LastName:  "González",
			Address:   stringPtr("Calle Belgrano 456"),
			Cellphone: stringPtr("11-3456-7890"),
			Email:     "2@mail.com",
			Username:  "mgonzalez",
			Password:  "2",
			IsActive:  true,
			IsAdmin:   false,
			RoleID:    2,
		},
		{
			ID:        4,
			FirstName: "Roberto",
			LastName:  "Martínez",
			Address:   stringPtr("Pasaje Los Álamos 789"),
			Cellphone: stringPtr("11-4567-8901"),
			Email:     "3@mail.com",
			Username:  "rmartinez",
			Password:  "3",
			IsActive:  true,
			IsAdmin:   false,
			RoleID:    3,
		},
	}

	for _, user := range users {
		dbTest.FirstOrCreate(&user, models.User{ID: user.ID})
	}

	// 6. Crear 20 Productos
	products := []models.Product{
		{ID: 1, Code: "BEB-001", Name: "Coca Cola 500ml", Description: stringPtr("Bebida gaseosa de cola"), Price: 250.0, CategoryID: 1, Notifier: true, MinAmount: 20.0},
		{ID: 2, Code: "BEB-002", Name: "Agua Mineral 500ml", Description: stringPtr("Agua mineral sin gas"), Price: 150.0, CategoryID: 1, Notifier: false, MinAmount: 10.0},
		{ID: 3, Code: "BEB-003", Name: "Powerade 500ml", Description: stringPtr("Bebida deportiva isotónica"), Price: 300.0, CategoryID: 1, Notifier: false, MinAmount: 10.0},
		{ID: 4, Code: "BEB-004", Name: "Jugo de Naranja 300ml", Description: stringPtr("Jugo natural de naranja"), Price: 200.0, CategoryID: 1, Notifier: true, MinAmount: 10.0},
		{ID: 5, Code: "SNK-001", Name: "Papas Fritas", Description: stringPtr("Papas fritas sabor original"), Price: 180.0, CategoryID: 2, Notifier: true, MinAmount: 10.0},
		{ID: 6, Code: "SNK-002", Name: "Maní Salado", Description: stringPtr("Maní tostado y salado"), Price: 120.0, CategoryID: 2, Notifier: true, MinAmount: 10.0},
		{ID: 7, Code: "SNK-003", Name: "Chocolate Barritas", Description: stringPtr("Barrita de chocolate con leche"), Price: 90.0, CategoryID: 2, Notifier: false, MinAmount: 10.0},
		{ID: 8, Code: "SNK-004", Name: "Galletas Saladas", Description: stringPtr("Galletas crackers saladas"), Price: 110.0, CategoryID: 2, Notifier: true, MinAmount: 10.0},
		{ID: 9, Code: "EQP-001", Name: "Pelota de Fútbol", Description: stringPtr("Pelota de fútbol profesional"), Price: 2500.0, CategoryID: 3, Notifier: false, MinAmount: 10.0},
		{ID: 10, Code: "EQP-002", Name: "Conos de Entrenamiento", Description: stringPtr("Set de 6 conos para entrenamientos"), Price: 800.0, CategoryID: 3, Notifier: true, MinAmount: 10.0},
		{ID: 11, Code: "EQP-003", Name: "Red de Fútbol", Description: stringPtr("Red para arco de fútbol 11"), Price: 1200.0, CategoryID: 3, Notifier: false, MinAmount: 10.0},
		{ID: 12, Code: "EQP-004", Name: "Pelota de Básquet", Description: stringPtr("Pelota de básquetbol oficial"), Price: 1800.0, CategoryID: 3, Notifier: true, MinAmount: 10.0},
		{ID: 13, Code: "ACC-001", Name: "Toalla Deportiva", Description: stringPtr("Toalla de microfibra"), Price: 450.0, CategoryID: 4, Notifier: true, MinAmount: 10.0},
		{ID: 14, Code: "ACC-002", Name: "Botella Deportiva", Description: stringPtr("Botella para hidratación 750ml"), Price: 320.0, CategoryID: 4, Notifier: true, MinAmount: 10.0},
		{ID: 15, Code: "ACC-003", Name: "Muñequeras", Description: stringPtr("Par de muñequeras deportivas"), Price: 280.0, CategoryID: 4, Notifier: false, MinAmount: 10.0},
		{ID: 16, Code: "COM-001", Name: "Sandwich de Jamón y Queso", Description: stringPtr("Sandwich casero"), Price: 400.0, CategoryID: 5, Notifier: true, MinAmount: 10.0},
		{ID: 17, Code: "COM-002", Name: "Empanadas", Description: stringPtr("Empanadas de carne (por unidad)"), Price: 150.0, CategoryID: 5, Notifier: true, MinAmount: 10.0},
		{ID: 18, Code: "COM-003", Name: "Pizza Slice", Description: stringPtr("Porción de pizza muzzarella"), Price: 350.0, CategoryID: 5, Notifier: true, MinAmount: 10.0},
		{ID: 19, Code: "BEB-005", Name: "Café Express", Description: stringPtr("Café espresso caliente"), Price: 180.0, CategoryID: 1, Notifier: true, MinAmount: 10.0},
		{ID: 20, Code: "SNK-005", Name: "Alfajor Triple", Description: stringPtr("Alfajor de dulce de leche"), Price: 220.0, CategoryID: 2, Notifier: true, MinAmount: 10.0},
	}

	for _, product := range products {
		dbTest.FirstOrCreate(&product, models.Product{ID: product.ID})
	}

	// 7. Crear Stock en Depósito
	stocksDeposit := []models.StockDeposit{
		{ProductID: 1, Stock: 100.0},
		{ProductID: 2, Stock: 150.0},
		{ProductID: 3, Stock: 80.0},
		{ProductID: 4, Stock: 60.0},
		{ProductID: 5, Stock: 200.0},
		{ProductID: 6, Stock: 120.0},
		{ProductID: 7, Stock: 300.0},
		{ProductID: 8, Stock: 180.0},
		{ProductID: 9, Stock: 10.0},
		{ProductID: 10, Stock: 15.0},
		{ProductID: 11, Stock: 5.0},
		{ProductID: 12, Stock: 8.0},
		{ProductID: 13, Stock: 50.0},
		{ProductID: 14, Stock: 40.0},
		{ProductID: 15, Stock: 60.0},
		{ProductID: 16, Stock: 20.0},
		{ProductID: 17, Stock: 100.0},
		{ProductID: 18, Stock: 30.0},
		{ProductID: 19, Stock: 200.0},
		{ProductID: 20, Stock: 150.0},
	}

	for _, stock := range stocksDeposit {
		dbTest.FirstOrCreate(&stock, models.StockDeposit{ProductID: stock.ProductID})
	}

	movementStock := []models.MovementStock{
		{UserID: 4, ProductID: 1, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 2, Amount: 75.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 3, Amount: 40.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 5, Amount: 100.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 7, Amount: 150.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 16, Amount: 10.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 17, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},

		{UserID: 4, ProductID: 1, Amount: 30.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 2, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 3, Amount: 25.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 13, Amount: 20.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 14, Amount: 15.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},

		{UserID: 4, ProductID: 4, Amount: 20.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 6, Amount: 40.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 8, Amount: 60.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 19, Amount: 80.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
		{UserID: 4, ProductID: 20, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
	}

	for _, movement := range movementStock {
		dbTest.FirstOrCreate(&movement, models.MovementStock{
			UserID:      movement.UserID,
			ProductID:   movement.ProductID,
			Amount:      movement.Amount,
			FromID:      movement.FromID,
			FromType:    movement.FromType,
			ToID:        movement.ToID,
			ToType:      movement.ToType,
			IgnoreStock: movement.IgnoreStock,
		})
	}

	// 8. Crear Stock en Puntos de Venta
	stocksPointSale := []models.StockPointSale{
		// Punto Central (ID: 1)
		{ProductID: 1, PointSaleID: 1, Stock: 50.0},
		{ProductID: 2, PointSaleID: 1, Stock: 75.0},
		{ProductID: 3, PointSaleID: 1, Stock: 40.0},
		{ProductID: 5, PointSaleID: 1, Stock: 100.0},
		{ProductID: 7, PointSaleID: 1, Stock: 150.0},
		{ProductID: 16, PointSaleID: 1, Stock: 10.0},
		{ProductID: 17, PointSaleID: 1, Stock: 50.0},

		// Cancha Norte (ID: 2)
		{ProductID: 1, PointSaleID: 2, Stock: 30.0},
		{ProductID: 2, PointSaleID: 2, Stock: 50.0},
		{ProductID: 3, PointSaleID: 2, Stock: 25.0},
		{ProductID: 13, PointSaleID: 2, Stock: 20.0},
		{ProductID: 14, PointSaleID: 2, Stock: 15.0},

		// Área de Descanso (ID: 3)
		{ProductID: 4, PointSaleID: 3, Stock: 20.0},
		{ProductID: 6, PointSaleID: 3, Stock: 40.0},
		{ProductID: 8, PointSaleID: 3, Stock: 60.0},
		{ProductID: 19, PointSaleID: 3, Stock: 80.0},
		{ProductID: 20, PointSaleID: 3, Stock: 50.0},
	}

	for _, stock := range stocksPointSale {
		dbTest.FirstOrCreate(&stock, models.StockPointSale{ProductID: stock.ProductID, PointSaleID: stock.PointSaleID})
	}

	// 9. Crear relaciones many-to-many

	// Usuarios - Puntos de Venta
	var user1, user2, user3 models.User
	var point1, point2, point3 models.PointSale

	dbTest.First(&user1, 2)
	dbTest.First(&user2, 3)
	dbTest.First(&user3, 4)
	dbTest.First(&point1, 1)
	dbTest.First(&point2, 2)
	dbTest.First(&point3, 3)

	// Asignar usuarios a puntos de venta
	dbTest.Model(&user1).Association("PointSales").Append([]models.PointSale{point1, point2, point3})
	dbTest.Model(&user2).Association("PointSales").Append([]models.PointSale{point1, point2})
	dbTest.Model(&user3).Association("PointSales").Append([]models.PointSale{point3})

	// Puntos de Venta - Canchas Deportivas
	var court1, court2, court3, court4 models.SportsCourt
	dbTest.First(&court1, 1)
	dbTest.First(&court2, 2)
	dbTest.First(&court3, 3)
	dbTest.First(&court4, 4)

	mountClose := 20000.0
	now := time.Now().UTC()

	register := []models.Register{
		{ID: 1, PointSaleID: point1.ID, UserOpenID: users[2].ID, OpenAmount: 10000.0, HourOpen: now, UserCloseID: &users[2].ID, CloseAmount: &mountClose, HourClose: &now, IsClose: true},
		{ID: 2, PointSaleID: point2.ID, UserOpenID: users[2].ID, OpenAmount: 15000.0, HourOpen: now, UserCloseID: nil, CloseAmount: nil, HourClose: nil, IsClose: false},
		{ID: 3, PointSaleID: point3.ID, UserOpenID: users[2].ID, OpenAmount: 20000.0, HourOpen: now, UserCloseID: nil, CloseAmount: nil, HourClose: nil, IsClose: false},
	}

	for _, reg := range register {
		dbTest.FirstOrCreate(&reg, models.Register{PointSaleID: reg.PointSaleID})
	}

	var register1, register2, register3 models.Register
	dbTest.First(&register1, 1)
	dbTest.First(&register2, 2)
	dbTest.First(&register3, 3)

	items := []models.IncomeItem{
		{ID: 1, IncomeID: 1, ProductID: products[0].ID, Quantity: 2, Price: products[0].Price, Subtotal: products[0].Price * 2, Price_Cost: products[0].Price},
		{ID: 2, IncomeID: 1, ProductID: products[1].ID, Quantity: 4, Price: products[1].Price, Subtotal: products[0].Price * 4, Price_Cost: products[1].Price},
		{ID: 3, IncomeID: 1, ProductID: products[2].ID, Quantity: 6, Price: products[2].Price, Subtotal: products[0].Price * 6, Price_Cost: products[2].Price},

		{ID: 4, IncomeID: 2, ProductID: products[0].ID, Quantity: 3, Price: products[0].Price, Subtotal: products[0].Price * 3, Price_Cost: products[0].Price},

		{ID: 5, IncomeID: 3, ProductID: products[1].ID, Quantity: 5, Price: products[1].Price, Subtotal: products[0].Price * 5, Price_Cost: products[1].Price},
		{ID: 6, IncomeID: 3, ProductID: products[2].ID, Quantity: 7, Price: products[2].Price, Subtotal: products[0].Price * 7, Price_Cost: products[2].Price},

		{ID: 7, IncomeID: 4, ProductID: products[0].ID, Quantity: 2, Price: products[0].Price, Subtotal: products[0].Price * 2, Price_Cost: products[0].Price},
		{ID: 8, IncomeID: 4, ProductID: products[1].ID, Quantity: 2, Price: products[1].Price, Subtotal: products[0].Price * 2, Price_Cost: products[1].Price},

		{ID: 9, IncomeID: 5, ProductID: products[2].ID, Quantity: 1, Price: products[2].Price, Subtotal: products[0].Price * 1, Price_Cost: products[2].Price},
		{ID: 10, IncomeID: 5, ProductID: products[0].ID, Quantity: 5, Price: products[0].Price, Subtotal: products[0].Price * 5, Price_Cost: products[0].Price},

		{ID: 11, IncomeID: 6, ProductID: products[1].ID, Quantity: 6, Price: products[1].Price, Subtotal: products[0].Price * 6, Price_Cost: products[1].Price},
		{ID: 12, IncomeID: 6, ProductID: products[2].ID, Quantity: 4, Price: products[2].Price, Subtotal: products[0].Price * 4, Price_Cost: products[2].Price},
		{ID: 13, IncomeID: 6, ProductID: products[0].ID, Quantity: 6, Price: products[0].Price, Subtotal: products[0].Price * 6, Price_Cost: products[0].Price},
		{ID: 14, IncomeID: 6, ProductID: products[1].ID, Quantity: 3, Price: products[1].Price, Subtotal: products[0].Price * 3, Price_Cost: products[1].Price},

		{ID: 15, IncomeID: 7, ProductID: products[0].ID, Quantity: 6, Price: products[0].Price, Subtotal: products[0].Price * 6, Price_Cost: products[0].Price},

		{ID: 16, IncomeID: 8, ProductID: products[1].ID, Quantity: 9, Price: products[1].Price, Subtotal: products[0].Price * 9, Price_Cost: products[1].Price},
		{ID: 17, IncomeID: 8, ProductID: products[2].ID, Quantity: 7, Price: products[2].Price, Subtotal: products[0].Price * 7, Price_Cost: products[2].Price},

		{ID: 18, IncomeID: 9, ProductID: products[0].ID, Quantity: 1, Price: products[0].Price, Subtotal: products[0].Price * 1, Price_Cost: products[0].Price},
		{ID: 19, IncomeID: 9, ProductID: products[1].ID, Quantity: 1, Price: products[1].Price, Subtotal: products[0].Price * 1, Price_Cost: products[1].Price},
		{ID: 20, IncomeID: 9, ProductID: products[2].ID, Quantity: 4, Price: products[2].Price, Subtotal: products[0].Price * 4, Price_Cost: products[2].Price},
	}

	incomes := []models.Income{
		{ID: 1, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: items[0].Subtotal + items[1].Subtotal + items[2].Subtotal, PaymentMethod: "efectivo"},
		{ID: 2, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: items[3].Subtotal, PaymentMethod: "transferencia"},
		{ID: 3, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: items[4].Subtotal + items[5].Subtotal, PaymentMethod: "efectivo"},

		{ID: 4, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: items[6].Subtotal + items[7].Subtotal, PaymentMethod: "tarjeta"},
		{ID: 5, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: items[8].Subtotal + items[9].Subtotal, PaymentMethod: "efectivo"},
		{ID: 6, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: items[10].Subtotal + items[11].Subtotal + items[12].Subtotal + items[13].Subtotal, PaymentMethod: "tarjeta"},

		{ID: 7, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: items[14].Subtotal, PaymentMethod: "transferencia"},
		{ID: 8, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: items[15].Subtotal + items[16].Subtotal, PaymentMethod: "efectivo"},
		{ID: 9, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: items[17].Subtotal + items[18].Subtotal + items[19].Subtotal, PaymentMethod: "tarjeta"},
	}

	for _, income := range incomes {
		if err := dbTest.FirstOrCreate(&income).Error; err != nil {
			return err
		}
	}

	for _, item := range items {
		if err := dbTest.FirstOrCreate(&item).Error; err != nil {
			return err
		}
	}

	restPay:= 2000.0
	restMethodPay := "efectivo"
	IncomeSportsCourts := []models.IncomeSportsCourts{
		{ID: 1, PartialRegisterID: register1.ID, RestRegisterID: &register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Shift: "tarde", SportsCourtID: 1, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "efectivo", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 2, PartialRegisterID: register1.ID, RestRegisterID: &register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Shift: "tarde", SportsCourtID: 1, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "transferencia", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 3, PartialRegisterID: register1.ID, RestRegisterID: &register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Shift: "tarde", SportsCourtID: 1, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "efectivo", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},

		{ID: 4, PartialRegisterID: register2.ID, RestRegisterID: &register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Shift: "tarde", SportsCourtID: 2, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "tarjeta", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 5, PartialRegisterID: register2.ID, RestRegisterID: &register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Shift: "tarde", SportsCourtID: 2, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "efectivo", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 6, PartialRegisterID: register2.ID, RestRegisterID: &register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Shift: "tarde", SportsCourtID: 2, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "tarjeta", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},

		{ID: 7, PartialRegisterID: register3.ID, RestRegisterID: &register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Shift: "tarde", SportsCourtID: 4, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "transferencia", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 8, PartialRegisterID: register3.ID, RestRegisterID: &register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Shift: "tarde", SportsCourtID: 3, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "efectivo", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
		{ID: 9, PartialRegisterID: register3.ID, RestRegisterID: &register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Shift: "tarde", SportsCourtID: 4, DatePlay: now, PartialPay: 1000, RestPay: &restPay, Price: 3000, PartialPaymentMethod: "tarjeta", RestPaymentMethod: &restMethodPay, DatePartialPay: now, DateRestPay: &now},
	}

	for _, item := range IncomeSportsCourts {
		if err := dbTest.FirstOrCreate(&item).Error; err != nil {
			return err
		}
	}

	expenses := []models.Expense{
		{ID: 1, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: 1000.0, PaymentMethod: "efectivo"},
		{ID: 2, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: 2000.0, PaymentMethod: "transferencia"},
		{ID: 3, RegisterID: register1.ID, UserID: users[2].ID, PointSaleID: point1.ID, Description: nil, Total: 1500.0, PaymentMethod: "efectivo"},

		{ID: 4, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: 2500.0, PaymentMethod: "tarjeta"},
		{ID: 5, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: 3000.0, PaymentMethod: "efectivo"},
		{ID: 6, RegisterID: register2.ID, UserID: users[2].ID, PointSaleID: point2.ID, Description: nil, Total: 4000.0, PaymentMethod: "tarjeta"},

		{ID: 7, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: 1200.0, PaymentMethod: "transferencia"},
		{ID: 8, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: 1800.0, PaymentMethod: "efectivo"},
		{ID: 9, RegisterID: register3.ID, UserID: users[2].ID, PointSaleID: point3.ID, Description: nil, Total: 2400.0, PaymentMethod: "tarjeta"},
	}

	for _, expense := range expenses {
		if err := dbTest.FirstOrCreate(&expense).Error; err != nil {
			return err
		}
	}

	itemsExpenseBuy := []models.ItemExpenseBuy{
		{ID: 1, ExpenseBuyID: 1, ProductID: products[0].ID, Quantity: 2, Price: products[0].Price, Subtotal: products[0].Price * 2},
		{ID: 2, ExpenseBuyID: 1, ProductID: products[1].ID, Quantity: 4, Price: products[1].Price, Subtotal: products[0].Price * 4},
		{ID: 3, ExpenseBuyID: 1, ProductID: products[2].ID, Quantity: 6, Price: products[2].Price, Subtotal: products[0].Price * 6},

		{ID: 4, ExpenseBuyID: 2, ProductID: products[0].ID, Quantity: 3, Price: products[0].Price, Subtotal: products[0].Price * 3},

		{ID: 5, ExpenseBuyID: 3, ProductID: products[1].ID, Quantity: 5, Price: products[1].Price, Subtotal: products[0].Price * 5},
		{ID: 6, ExpenseBuyID: 3, ProductID: products[2].ID, Quantity: 7, Price: products[2].Price, Subtotal: products[0].Price * 7},

		{ID: 7, ExpenseBuyID: 4, ProductID: products[0].ID, Quantity: 2, Price: products[0].Price, Subtotal: products[0].Price * 2},
		{ID: 8, ExpenseBuyID: 4, ProductID: products[1].ID, Quantity: 2, Price: products[1].Price, Subtotal: products[0].Price * 2},

		{ID: 9, ExpenseBuyID: 5, ProductID: products[2].ID, Quantity: 1, Price: products[2].Price, Subtotal: products[0].Price * 1},
		{ID: 10, ExpenseBuyID: 5, ProductID: products[0].ID, Quantity: 5, Price: products[0].Price, Subtotal: products[0].Price * 5},

		{ID: 11, ExpenseBuyID: 6, ProductID: products[1].ID, Quantity: 6, Price: products[1].Price, Subtotal: products[0].Price * 6},
		{ID: 12, ExpenseBuyID: 6, ProductID: products[2].ID, Quantity: 4, Price: products[2].Price, Subtotal: products[0].Price * 4},
		{ID: 13, ExpenseBuyID: 6, ProductID: products[0].ID, Quantity: 6, Price: products[0].Price, Subtotal: products[0].Price * 6},
		{ID: 14, ExpenseBuyID: 6, ProductID: products[1].ID, Quantity: 3, Price: products[1].Price, Subtotal: products[0].Price * 3},

		{ID: 15, ExpenseBuyID: 7, ProductID: products[0].ID, Quantity: 6, Price: products[0].Price, Subtotal: products[0].Price * 6},

		{ID: 16, ExpenseBuyID: 8, ProductID: products[1].ID, Quantity: 9, Price: products[1].Price, Subtotal: products[0].Price * 9},
		{ID: 17, ExpenseBuyID: 8, ProductID: products[2].ID, Quantity: 7, Price: products[2].Price, Subtotal: products[0].Price * 7},

		{ID: 18, ExpenseBuyID: 9, ProductID: products[0].ID, Quantity: 1, Price: products[0].Price, Subtotal: products[0].Price * 1},
		{ID: 19, ExpenseBuyID: 9, ProductID: products[1].ID, Quantity: 1, Price: products[1].Price, Subtotal: products[0].Price * 1},
		{ID: 20, ExpenseBuyID: 9, ProductID: products[2].ID, Quantity: 4, Price: products[2].Price, Subtotal: products[0].Price * 4},
	}

	expenseBuy := []models.ExpenseBuy{
		{ID: 1, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[0].Subtotal + itemsExpenseBuy[1].Subtotal + itemsExpenseBuy[2].Subtotal, PaymentMethod: "efectivo"},
		{ID: 2, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[3].Subtotal, PaymentMethod: "transferencia"},
		{ID: 3, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[4].Subtotal + itemsExpenseBuy[5].Subtotal, PaymentMethod: "efectivo"},

		{ID: 4, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[6].Subtotal + itemsExpenseBuy[7].Subtotal, PaymentMethod: "tarjeta"},
		{ID: 5, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[8].Subtotal + itemsExpenseBuy[9].Subtotal, PaymentMethod: "efectivo"},
		{ID: 6, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[10].Subtotal + itemsExpenseBuy[11].Subtotal + itemsExpenseBuy[12].Subtotal + itemsExpenseBuy[13].Subtotal, PaymentMethod: "tarjeta"},

		{ID: 7, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[14].Subtotal, PaymentMethod: "transferencia"},
		{ID: 8, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[15].Subtotal + itemsExpenseBuy[16].Subtotal, PaymentMethod: "efectivo"},
		{ID: 9, UserID: users[1].ID, Description: nil, Total: itemsExpenseBuy[17].Subtotal + itemsExpenseBuy[18].Subtotal + itemsExpenseBuy[19].Subtotal, PaymentMethod: "tarjeta"},
	}

	for _, expense := range expenseBuy {
		if err := dbTest.FirstOrCreate(&expense).Error; err != nil {
			return err
		}
	}

	for _, item := range itemsExpenseBuy {
		if err := dbTest.FirstOrCreate(&item).Error; err != nil {
			return err
		}
	}

	return nil
}

func DeleteTestData() error {
	fmt.Println("Iniciando eliminación de datos...")

	// Orden de eliminación respetando foreign keys
	queries := []string{
		// 1. Eliminar relaciones many-to-many primero
		"DELETE FROM user_point_sales",
		"DELETE FROM income_sports_courts",
		"DELETE FROM sports_courts_point_sales",
		"DELETE FROM movement_stocks",
		"DELETE FROM income_items",
		"DELETE FROM item_expense_buys",

		// 2. Eliminar tablas con foreign keys (orden de dependencias)
		"DELETE FROM stock_point_sales",
		"DELETE FROM stock_deposits",
		"DELETE FROM products",
		"DELETE FROM incomes",
		"DELETE FROM expenses",
		"DELETE FROM expense_buys",
		"DELETE FROM registers",

		// 3. Eliminar tablas independientes
		"DELETE FROM sports_courts",
		"DELETE FROM point_sales",
		"DELETE FROM categories",
	}

	// Ejecutar cada query
	for i, query := range queries {
		fmt.Printf("Ejecutando paso %d: %s\n", i+1, query)

		if err := dbTest.Exec(query).Error; err != nil {
			return fmt.Errorf("error ejecutando '%s': %v", query, err)
		}
	}

	fmt.Println("✓ Todos los datos eliminados exitosamente!")
	return nil
}

func stringPtr(s string) *string {
	return &s
}
