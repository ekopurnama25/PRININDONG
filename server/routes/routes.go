package routes

import (
	middlewares "apk-chat-serve/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	SetupPublicAuthRoutes(app)
	app.Use(middlewares.IsUserAuthenticated)
	SetupAuthRoutes(app)
	SetupSaldoRoutes(app)
	SetupUserRoutes(app)
	SetupRolesRoutes(app)
	SetupTintaRoutes(app)
	SetupPesananRoutes(app)
	SetupConfirmationRoutes(app)
}
