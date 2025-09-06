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
		&models.Expense{}, &models.Income{}, &models.IncomeSportsCourts{},
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
