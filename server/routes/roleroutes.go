package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRolesRoutes(app *fiber.App) {
	app.Get("/api/roles/", controllers.AllRoles)
	app.Post("/api/roles/", controllers.SaveRoles)
	app.Get("/api/roles/:id", controllers.GetIdRoles)
	app.Delete("/api/roles/:id", controllers.DeleteIdRole)
	app.Put("/api/roles/:id", controllers.UpdateRoles)
}
