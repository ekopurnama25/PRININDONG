package routes

import (
	"apk-chat-serve/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicAuthRoutes(app *fiber.App) {
	app.Post("/api/refreshToken/", controllers.PostRefreshToken)
	app.Post("/api/login/", controllers.AUthUsersMiddlaware)
}

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/api/home/", controllers.GetUsersLogin)
	app.Post("/api/logout/", controllers.LogoutAuth)
}
