package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/api/users/", controllers.AllUsers)
	app.Post("/api/users/", controllers.SaveUsers)
	app.Get("/api/users/:id", controllers.GetIdUsers)
	app.Delete("/api/users/:id", controllers.DeleteIdUsers)
	app.Put("/api/users/:id", controllers.UpdateUsers)
}
