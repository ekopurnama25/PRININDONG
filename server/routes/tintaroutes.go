package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupTintaRoutes(app *fiber.App) {
	app.Get("/api/tinta/", controllers.AllTinta)
	app.Get("/api/tinta/:id", controllers.GetIdTinta)
	app.Post("/api/tinta/", controllers.SaveTinta)
	app.Delete("/api/tinta/:id", controllers.DeleteIdTinta)
	app.Put("/api/tinta/:id", controllers.UpdateTinta)
}
