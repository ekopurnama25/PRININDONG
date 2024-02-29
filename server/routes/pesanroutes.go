package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPesananRoutes(app *fiber.App) {
	app.Post("/api/pesanan/", controllers.SavePesanan)
}
