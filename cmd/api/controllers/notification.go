package controllers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

//	 NotificationAlert godoc
//	@Summary		NotificationAlert
//	@Description	Enviar notificaciones
//	@Tags			Notification
//	@Accept			json
//	@Produce		text/event-stream
//	@Security		CookieAuth
//	@Router			/api/v1/notification/alert [get]
func (n *NotificationController) NotificationAlert(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	ctx := c.Context()

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		// Enviar notificación inicial
		n.sendStockNotificationSSE(w, ctx)

		for {
			select {
			case <-ctx.Done():
				log.Println("Cliente desconectado de SSE")
				return
			case <-n.NotifyCh: // 🔥 se dispara desde otro endpoint
				if err := n.sendStockNotificationSSE(w, ctx); err != nil {
					log.Printf("Error enviando notificación SSE: %v", err)
					return
				}
			}
		}
	})

	return nil
}

func (n *NotificationController) sendStockNotificationSSE(w *bufio.Writer, ctx context.Context) error {
	// Verificar si el contexto fue cancelado
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	products, err := n.NotificationService.NotificationStock()
	if err != nil {
		log.Printf("Error obteniendo productos con stock bajo: %v", err)
		// Enviar evento de error
		fmt.Fprintf(w, "event: error\n")
		fmt.Fprintf(w, "data: {\"error\": \"Error obteniendo productos\"}\n\n")
		w.Flush()
		return err
	}

	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")

	// Convertir a JSON
	data, err := json.Marshal(schemas.Response{
		Status: true,
		Message: "Productos con stock bajo",
		Body: map[string]any{
			"event":    "alert-stock",
			"response": map[string]any{
				"products": products,
				"count":    len(products),
				"datetime": time.Now().In(loc),
			},
		},
	},
	)
	// data, err := json.Marshal(map[string]any{
	// 	"event":     "alert-stock",
	// 	"products":  products,
	// 	"count":     len(products),
	// 	"datetime": time.Now().In(loc),
	// })

	if err != nil {
		log.Printf("Error serializando datos: %v", err)
		return err
	}

	// Enviar evento SSE
	fmt.Fprintf(w, "event: stock-notification\n")
	fmt.Fprintf(w, "data: %s\n\n", string(data))
	w.Flush()

	return nil
}
