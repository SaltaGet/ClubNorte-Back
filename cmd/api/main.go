package main

import (
	"log"
	"os"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/jobs"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/routes"
	_ "github.com/DanielChachagua/Club-Norte-Back/cmd/api/docs"
	"github.com/DanielChachagua/Club-Norte-Back/internal/database"
	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

//	@title						APP Club Norte
//	@version					1.0
//	@description				This is a API for Club Norte
//	@termsOfService				http://swagger.io/terms/
//	@securityDefinitions.apikey	CookieAuth
//	@in							cookie
//	@name						access_token
//	@description				Enter your JWT token here. Example: "eyJhbGciOiJIUz..."


func main() {
	err := godotenv.Load("/etc/variables/club-norte/.env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	local := os.Getenv("LOCAL")
	if local == "true" {
		if err := jobs.GenerateSwagger(); err != nil {
			log.Fatalf("Error ejecutando swag init: %v", err)
		}
	}

	dep := dependencies.NewMainContainer(db)

	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})

	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.InjectionDepends(dep))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        120,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Demasiadas peticiones. Intentá más tarde.",
			})
		},
	}))

	routes.SetupRoutes(app, dep)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
