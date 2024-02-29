package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupConfirmationRoutes(app *fiber.App) {
	app.Post("/api/confirmation/users/:id", controllers.ConfirmasiUsersPesananUsers)
	app.Post("/api/confirmation/admin/:id", controllers.ConfirmasiAdminPesananUsers)
	app.Post("/api/confirmationselesai/admin/:id", controllers.ConfirmasiStatusSelesai)
}
