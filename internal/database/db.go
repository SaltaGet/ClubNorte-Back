package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// "github.com/google/uuid"
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbTest *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	env := os.Getenv("ENV")
	if env == "prod" {
		return connectMySQL()
	}
	return connectSQLite()
}

func connectSQLite() (*gorm.DB, error) {
	uri := os.Getenv("URI_DB_DEV")
	if uri == "" {
		return nil, fmt.Errorf("la variable de entorno URI_DB_DEV no esta definida")
	}

	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	setupDBConnection(db, 10, 5)

	if err := db.AutoMigrate(
		&models.User{}, &models.Category{},
		&models.Expense{}, &models.Income{}, &models.IncomeSportsCourts{},
		&models.MovementStock{}, &models.PointSale{}, &models.Product{},
		&models.PointSale{}, &models.Register{}, &models.Role{},
		&models.SportsCourt{}, &models.StockDeposit{}, &models.StockPointSale{},
	); err != nil {
		return nil, err
	}

	err = initialData(db)
	if err != nil {
		return nil, err
	}

	dbTest = db

	return ensureAdmin(db)
}

// func connectMySQL() (*gorm.DB, error) {
// 	dsn := os.Getenv("URI_DB_PROD")
// 	if dsn == "" {
// 		return nil, fmt.Errorf("la variable de entorno URI_DB_PROD no esta definida")
// 	}

// 	if err := ensureDatabaseExists(dsn); err != nil {
// 		log.Fatalf("No se pudo crear la base: %v", err)
// 	}

// 	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	var db *gorm.DB
// 	var err error

// 	// Intentar varias veces hasta que MySQL esté listo
// 	for i := 0; i < 10; i++ {
// 		if err := ensureDatabaseExists(dsn); err != nil {
// 			log.Printf("Esperando a que MySQL esté listo... intento %d: %v", i+1, err)
// 			time.Sleep(3 * time.Second)
// 			continue
// 		}

// 		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err == nil {
// 			break
// 		}

// 		log.Printf("Error conectando a MySQL, retry %d: %v", i+1, err)
// 		time.Sleep(3 * time.Second)
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("no se pudo conectar a MySQL después de varios intentos: %v", err)
// 	}

// 	setupDBConnection(db, 15, 5)

// 	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
// 		&models.User{}, &models.Category{},
// 		&models.Expense{}, &models.Income{}, &models.IncomeSportsCourts{},
// 		&models.MovementStock{}, &models.PointSale{}, &models.Product{},
// 		&models.PointSale{}, &models.Register{}, &models.Role{},
// 		&models.SportsCourt{}, &models.StockDeposit{}, &models.StockPointSale{},
// 	); err != nil {
// 		log.Fatalf("Error en migración: %v", err)
// 	}

// 	err = initialData(db)
// 	if err != nil {
// 		return nil, err
// 	}

//		return ensureAdmin(db)
//	}
func connectMySQL() (*gorm.DB, error) {
	dsn := os.Getenv("URI_DB_PROD")
	if dsn == "" {
		return nil, fmt.Errorf("la variable de entorno URI_DB_PROD no esta definida")
	}

	var db *gorm.DB
	var err error

	// Intentar varias veces hasta que MariaDB esté listo (más intentos)
	maxRetries := 20
	if ensureErr := ensureDatabaseExists(dsn); ensureErr != nil {
		log.Printf("Warning: Error creando base de datos: %v", ensureErr)
		// No fatal, continuar con la conexión existente
	}

	for i := 0; i < maxRetries; i++ {
		// Primero intentar conectar directamente
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Verificar si la conexión es realmente válida
			sqlDB, pingErr := db.DB()
			if pingErr == nil {
				pingErr = sqlDB.Ping()
				if pingErr == nil {
					log.Println("✅ Conexión a la base de datos exitosa")

					// Ahora asegurar que la base de datos existe

					break // Salir del loop si todo está bien
				}
			}
		}

		log.Printf("Esperando a que MariaDB esté listo... intento %d/%d: %v",
			i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a MySQL después de %d intentos: %v", maxRetries, err)
	}

	setupDBConnection(db, 15, 5)

	// Migraciones
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{}, &models.Category{},
		&models.Expense{}, &models.Income{}, &models.IncomeSportsCourts{}, &models.IncomeItem{},
		&models.MovementStock{}, &models.PointSale{}, &models.Product{},
		&models.PointSale{}, &models.Register{}, &models.Role{},
		&models.SportsCourt{}, &models.StockDeposit{}, &models.StockPointSale{},
	); err != nil {
		log.Printf("Error en migración: %v", err)
		// No fatal, continuar
	}

	err = initialData(db)
	if err != nil {
		log.Printf("Error en datos iniciales: %v", err)
		// No fatal, continuar
	}

	dbTest = db

	return ensureAdmin(db)
}

func ensureAdmin(db *gorm.DB) (*gorm.DB, error) {
	var email string
	db.Model(&models.User{}).Select("email").Where("email = ?", os.Getenv("ADMIN_EMAIL")).Scan(&email)

	if email != "" {
		log.Println("El admin ya existe")
		return db, nil
	}

	var role models.Role
	if err := db.Where("name = ?", os.Getenv("ADMIN_ROLE")).First(&role).Error; err != nil {
		return nil, fmt.Errorf("el rol %s no existe", os.Getenv("ADMIN_ROLE"))
	}

	address := os.Getenv("ADMIN_ADDRESS")
	cellphone := os.Getenv("ADMIN_CELLPHONE")

	if err := db.Create(&models.User{
		FirstName: os.Getenv("ADMIN_FIRST_NAME"),
		LastName:  os.Getenv("ADMIN_LAST_NAME"),
		Address:   &address,
		Cellphone: &cellphone,
		Email:     os.Getenv("ADMIN_EMAIL"),
		Username:  os.Getenv("ADMIN_USERNAME"),
		Password:  os.Getenv("ADMIN_PASSWORD"),
		RoleID:    role.ID,
		IsAdmin:   true,
	}).Error; err != nil {
		return nil, err
	}

	return db, nil
}

func setupDBConnection(db *gorm.DB, maxOpen, maxIdle int) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error al obtener conexión de base: %v", err)
	}
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
}

func ensureDatabaseExists(dsn string) error {
	if os.Getenv("ENV") != "prod" {
		return nil
	}

	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return fmt.Errorf("DSN inválido: %s", dsn)
	}
	dbNameAndParams := parts[1]
	dbName := strings.Split(dbNameAndParams, "?")[0]

	dsnWithoutDB := strings.Split(dsn, "/")[0] + "/?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return fmt.Errorf("error al conectar sin base: %w", err)
	}
	defer db.Close()

	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("no se pudo conectar al servidor MySQL: %w", err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName))
	return err
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("No se pudo obtener la conexión de bajo nivel:", err)
	}

	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Error al cerrar la conexión:", err)
		}
	}
	return nil
}


// func CreateTestData() error {
//     // 2. Crear Categorías
//     categories := []models.Category{
//         {ID: 1, Name: "Bebidas"},
//         {ID: 2, Name: "Snacks"},
//         {ID: 3, Name: "Equipamiento Deportivo"},
//         {ID: 4, Name: "Accesorios"},
//         {ID: 5, Name: "Comida"},
//     }
    
//     for _, category := range categories {
//         dbTest.FirstOrCreate(&category, models.Category{ID: category.ID})
//     }

//     // 3. Crear Puntos de Venta
//     pointSales := []models.PointSale{
//         {
//             ID: 1, 
//             Name: "Punto Central", 
//             Description: stringPtr("Punto de venta principal del complejo"),
//         },
//         {
//             ID: 2, 
//             Name: "Cancha Norte", 
//             Description: stringPtr("Kiosco ubicado en las canchas del sector norte"),
//         },
//         {
//             ID: 3, 
//             Name: "Área de Descanso", 
//             Description: stringPtr("Punto de venta en la zona de descanso"),
//         },
//     }
    
//     for _, pointSale := range pointSales {
//         dbTest.FirstOrCreate(&pointSale, models.PointSale{ID: pointSale.ID})
//     }

//     // 4. Crear Canchas Deportivas
//     sportsCourts := []models.SportsCourt{
//         {
//             ID: 1,
//             Code: "CANCHA-01",
//             Name: "Cancha Fútbol A",
//             Description: stringPtr("Cancha de fútbol principal con césped sintético"),
//         },
//         {
//             ID: 2,
//             Code: "CANCHA-02",
//             Name: "Cancha Fútbol B",
//             Description: stringPtr("Cancha de fútbol secundaria"),
//         },
//         {
//             ID: 3,
//             Code: "CANCHA-03",
//             Name: "Cancha Tenis 1",
//             Description: stringPtr("Cancha de tenis con superficie de polvo de ladrillo"),
//         },
//         {
//             ID: 4,
//             Code: "CANCHA-04",
//             Name: "Cancha Básquet",
//             Description: stringPtr("Cancha de básquetbol techada"),
//         },
//     }
    
//     for _, court := range sportsCourts {
//         dbTest.FirstOrCreate(&court, models.SportsCourt{ID: court.ID})
//     }

// 		var sportsCourt1, sportsCourt2, sportsCourt3, sportsCourt4 models.SportsCourt
//     var pointSport1, pointSport2, pointSport3 models.PointSale
    
//     dbTest.First(&sportsCourt1, 1)
//     dbTest.First(&sportsCourt2, 2)
//     dbTest.First(&sportsCourt3, 3)
//     dbTest.First(&sportsCourt4, 4)
//     dbTest.First(&pointSport1, 1)
//     dbTest.First(&pointSport2, 2)
//     dbTest.First(&pointSport3, 3)
    
//     // Asignar usuarios a puntos de venta
//     dbTest.Model(&pointSport1).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt1})
//     dbTest.Model(&pointSport2).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt2})
//     dbTest.Model(&pointSport3).Association("SportsCourts").Append([]models.SportsCourt{sportsCourt3, sportsCourt4})

//     // 5. Crear Usuarios
//     users := []models.User{
//         {
//             ID: 2,
//             FirstName: "Juan Carlos",
//             LastName: "Pérez",
//             Address: stringPtr("Av. San Martín 123"),
//             Cellphone: stringPtr("11-2345-6789"),
//             Email: "1@mail.com",
//             Username: "jperez",
//             Password: "1", // En producción usar hash real
//             IsActive: true,
//             IsAdmin: false,
//             RoleID: 1,
//         },
//         {
//             ID: 3,
//             FirstName: "María Elena",
//             LastName: "González",
//             Address: stringPtr("Calle Belgrano 456"),
//             Cellphone: stringPtr("11-3456-7890"),
//             Email: "2@mail.com",
//             Username: "mgonzalez",
//             Password: "2",
//             IsActive: true,
//             IsAdmin: false,
//             RoleID: 2,
//         },
//         {
//             ID: 4,
//             FirstName: "Roberto",
//             LastName: "Martínez",
//             Address: stringPtr("Pasaje Los Álamos 789"),
//             Cellphone: stringPtr("11-4567-8901"),
//             Email: "3@mail.com",
//             Username: "rmartinez",
//             Password: "3",
//             IsActive: true,
//             IsAdmin: false,
//             RoleID: 3,
//         },
//     }
    
//     for _, user := range users {
//         dbTest.FirstOrCreate(&user, models.User{ID: user.ID})
//     }

//     // 6. Crear 20 Productos
//     products := []models.Product{
//         {ID: 1, Code: "BEB-001", Name: "Coca Cola 500ml", Description: stringPtr("Bebida gaseosa de cola"), Price: 250.0, CategoryID: 1},
//         {ID: 2, Code: "BEB-002", Name: "Agua Mineral 500ml", Description: stringPtr("Agua mineral sin gas"), Price: 150.0, CategoryID: 1},
//         {ID: 3, Code: "BEB-003", Name: "Powerade 500ml", Description: stringPtr("Bebida deportiva isotónica"), Price: 300.0, CategoryID: 1},
//         {ID: 4, Code: "BEB-004", Name: "Jugo de Naranja 300ml", Description: stringPtr("Jugo natural de naranja"), Price: 200.0, CategoryID: 1},
//         {ID: 5, Code: "SNK-001", Name: "Papas Fritas", Description: stringPtr("Papas fritas sabor original"), Price: 180.0, CategoryID: 2},
//         {ID: 6, Code: "SNK-002", Name: "Maní Salado", Description: stringPtr("Maní tostado y salado"), Price: 120.0, CategoryID: 2},
//         {ID: 7, Code: "SNK-003", Name: "Chocolate Barritas", Description: stringPtr("Barrita de chocolate con leche"), Price: 90.0, CategoryID: 2},
//         {ID: 8, Code: "SNK-004", Name: "Galletas Saladas", Description: stringPtr("Galletas crackers saladas"), Price: 110.0, CategoryID: 2},
//         {ID: 9, Code: "EQP-001", Name: "Pelota de Fútbol", Description: stringPtr("Pelota de fútbol profesional"), Price: 2500.0, CategoryID: 3},
//         {ID: 10, Code: "EQP-002", Name: "Conos de Entrenamiento", Description: stringPtr("Set de 6 conos para entrenamientos"), Price: 800.0, CategoryID: 3},
//         {ID: 11, Code: "EQP-003", Name: "Red de Fútbol", Description: stringPtr("Red para arco de fútbol 11"), Price: 1200.0, CategoryID: 3},
//         {ID: 12, Code: "EQP-004", Name: "Pelota de Básquet", Description: stringPtr("Pelota de básquetbol oficial"), Price: 1800.0, CategoryID: 3},
//         {ID: 13, Code: "ACC-001", Name: "Toalla Deportiva", Description: stringPtr("Toalla de microfibra"), Price: 450.0, CategoryID: 4},
//         {ID: 14, Code: "ACC-002", Name: "Botella Deportiva", Description: stringPtr("Botella para hidratación 750ml"), Price: 320.0, CategoryID: 4},
//         {ID: 15, Code: "ACC-003", Name: "Muñequeras", Description: stringPtr("Par de muñequeras deportivas"), Price: 280.0, CategoryID: 4},
//         {ID: 16, Code: "COM-001", Name: "Sandwich de Jamón y Queso", Description: stringPtr("Sandwich casero"), Price: 400.0, CategoryID: 5},
//         {ID: 17, Code: "COM-002", Name: "Empanadas", Description: stringPtr("Empanadas de carne (por unidad)"), Price: 150.0, CategoryID: 5},
//         {ID: 18, Code: "COM-003", Name: "Pizza Slice", Description: stringPtr("Porción de pizza muzzarella"), Price: 350.0, CategoryID: 5},
//         {ID: 19, Code: "BEB-005", Name: "Café Express", Description: stringPtr("Café espresso caliente"), Price: 180.0, CategoryID: 1},
//         {ID: 20, Code: "SNK-005", Name: "Alfajor Triple", Description: stringPtr("Alfajor de dulce de leche"), Price: 220.0, CategoryID: 2},
//     }
    
//     for _, product := range products {
//         dbTest.FirstOrCreate(&product, models.Product{ID: product.ID})
//     }

//     // 7. Crear Stock en Depósito
//     stocksDeposit := []models.StockDeposit{
//         {ProductID: 1, Stock: 100.0},
//         {ProductID: 2, Stock: 150.0},
//         {ProductID: 3, Stock: 80.0},
//         {ProductID: 4, Stock: 60.0},
//         {ProductID: 5, Stock: 200.0},
//         {ProductID: 6, Stock: 120.0},
//         {ProductID: 7, Stock: 300.0},
//         {ProductID: 8, Stock: 180.0},
//         {ProductID: 9, Stock: 10.0},
//         {ProductID: 10, Stock: 15.0},
//         {ProductID: 11, Stock: 5.0},
//         {ProductID: 12, Stock: 8.0},
//         {ProductID: 13, Stock: 50.0},
//         {ProductID: 14, Stock: 40.0},
//         {ProductID: 15, Stock: 60.0},
//         {ProductID: 16, Stock: 20.0},
//         {ProductID: 17, Stock: 100.0},
//         {ProductID: 18, Stock: 30.0},
//         {ProductID: 19, Stock: 200.0},
//         {ProductID: 20, Stock: 150.0},
//     }
    
//     for _, stock := range stocksDeposit {
//         dbTest.FirstOrCreate(&stock, models.StockDeposit{ProductID: stock.ProductID})
//     }

// 		movementStock := []models.MovementStock{
// 			{UserID: 4, ProductID: 1, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 2, Amount: 75.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 3, Amount: 40.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 5, Amount: 100.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 7, Amount: 150.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 16, Amount: 10.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 17, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 1, ToType: "point_sale", IgnoreStock: false},

// 			{UserID: 4, ProductID: 1, Amount: 30.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 2, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 3, Amount: 25.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 13, Amount: 20.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 14, Amount: 15.0, FromID: 1, FromType: "deposit", ToID: 2, ToType: "point_sale", IgnoreStock: false},

// 			{UserID: 4, ProductID: 4, Amount: 20.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 6, Amount: 40.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 8, Amount: 60.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 19, Amount: 80.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
// 			{UserID: 4, ProductID: 20, Amount: 50.0, FromID: 1, FromType: "deposit", ToID: 3, ToType: "point_sale", IgnoreStock: false},
// 		}

// 		for _, movement := range movementStock {
//         dbTest.FirstOrCreate(&movement, models.MovementStock{
// 					UserID:      movement.UserID,
// 					ProductID:   movement.ProductID,
// 					Amount:      movement.Amount,
// 					FromID:      movement.FromID,
// 					FromType:    movement.FromType,
// 					ToID:        movement.ToID,
// 					ToType:      movement.ToType,
// 					IgnoreStock: movement.IgnoreStock,
// 				},)
//     }

//     // 8. Crear Stock en Puntos de Venta
//     stocksPointSale := []models.StockPointSale{
//         // Punto Central (ID: 1)
//         {ProductID: 1, PointSaleID: 1, Stock: 50.0},
//         {ProductID: 2, PointSaleID: 1, Stock: 75.0},
//         {ProductID: 3, PointSaleID: 1, Stock: 40.0},
//         {ProductID: 5, PointSaleID: 1, Stock: 100.0},
//         {ProductID: 7, PointSaleID: 1, Stock: 150.0},
//         {ProductID: 16, PointSaleID: 1, Stock: 10.0},
//         {ProductID: 17, PointSaleID: 1, Stock: 50.0},
        
//         // Cancha Norte (ID: 2)
//         {ProductID: 1, PointSaleID: 2, Stock: 30.0},
//         {ProductID: 2, PointSaleID: 2, Stock: 50.0},
//         {ProductID: 3, PointSaleID: 2, Stock: 25.0},
//         {ProductID: 13, PointSaleID: 2, Stock: 20.0},
//         {ProductID: 14, PointSaleID: 2, Stock: 15.0},
        
//         // Área de Descanso (ID: 3)
//         {ProductID: 4, PointSaleID: 3, Stock: 20.0},
//         {ProductID: 6, PointSaleID: 3, Stock: 40.0},
//         {ProductID: 8, PointSaleID: 3, Stock: 60.0},
//         {ProductID: 19, PointSaleID: 3, Stock: 80.0},
//         {ProductID: 20, PointSaleID: 3, Stock: 50.0},
//     }
    
//     for _, stock := range stocksPointSale {
//         dbTest.FirstOrCreate(&stock, models.StockPointSale{ProductID: stock.ProductID, PointSaleID: stock.PointSaleID})
//     }

//     // 9. Crear relaciones many-to-many
    
//     // Usuarios - Puntos de Venta
//     var user1, user2, user3 models.User
// 		var point1, point2, point3 models.PointSale
    
//     dbTest.First(&user1, 2)
//     dbTest.First(&user2, 3)
//     dbTest.First(&user3, 4)
//     dbTest.First(&point1, 1)
//     dbTest.First(&point2, 2)
//     dbTest.First(&point3, 3)
    
//     // Asignar usuarios a puntos de venta
//     dbTest.Model(&user1).Association("PointSales").Append([]models.PointSale{point1, point2, point3})
//     dbTest.Model(&user2).Association("PointSales").Append([]models.PointSale{point1, point2})
//     dbTest.Model(&user3).Association("PointSales").Append([]models.PointSale{point3})
    
//     // Puntos de Venta - Canchas Deportivas
//     var court1, court2, court3, court4 models.SportsCourt
//     dbTest.First(&court1, 1)
//     dbTest.First(&court2, 2)
//     dbTest.First(&court3, 3)
//     dbTest.First(&court4, 4)
    
//     // Nota: Hay un error en tu modelo SportsCourt, debería ser []PointSale en lugar de []User
//     // Asumo que quieres relacionar canchas con puntos de venta
//     // db.Model(&point1).Association("SportsCourts").Append([]SportsCourt{court1, court2})
//     // db.Model(&point2).Association("SportsCourts").Append([]SportsCourt{court1, court3, court4})
//     // db.Model(&point3).Association("SportsCourts").Append([]SportsCourt{court3})

//     return nil
// }

// func DeleteTestData() error {
// 	fmt.Println("Iniciando eliminación de datos...")
    
//     // Orden de eliminación respetando foreign keys
//     queries := []string{
//         // 1. Eliminar relaciones many-to-many primero
//         "DELETE FROM user_point_sales",
//         "DELETE FROM sports_courts_point_sales",
// 				"DELETE FROM movement_stocks",
        
//         // 2. Eliminar tablas con foreign keys (orden de dependencias)
//         "DELETE FROM stock_point_sales",
//         "DELETE FROM stock_deposits", 
//         "DELETE FROM products",
        
//         // 3. Eliminar tablas independientes
//         "DELETE FROM sports_courts",
//         "DELETE FROM point_sales", 
//         "DELETE FROM categories",
//     }
    
//     // Ejecutar cada query
//     for i, query := range queries {
//         fmt.Printf("Ejecutando paso %d: %s\n", i+1, query)
        
//         if err := dbTest.Exec(query).Error; err != nil {
//             return fmt.Errorf("error ejecutando '%s': %v", query, err)
//         }
//     }
    
//     fmt.Println("✓ Todos los datos eliminados exitosamente!")
//     return nil
// }


// func stringPtr(s string) *string {
//     return &s
// }