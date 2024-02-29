package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupSaldoRoutes(app *fiber.App) {
	app.Get("/api/saldo/", controllers.AllSaldo)
	app.Post("/api/saldo/", controllers.SaveSaldoUsers)
	app.Post("/api/topup/:id", controllers.IsiSaldoUsers)
	app.Get("/api/saldo/:id", controllers.GetIdSaldo)
	app.Delete("/api/saldo/:id", controllers.DeleteIdSaldo)
	app.Put("/api/saldo/:id", controllers.UpdateSaldoUsers)
}
