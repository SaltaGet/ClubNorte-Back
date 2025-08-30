package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now() 

	err := c.Next()

	duration := time.Since(start)

	c.Set("X-Response-Time", duration.String())

	log.Printf("Request: %s %s took %s", c.Method(), c.Path(), duration)

	return err
}