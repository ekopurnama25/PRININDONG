package main

import (
	"apk-chat-serve/config"
	"apk-chat-serve/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Static("/utils", "./utils")
	routes.Setup(app)

	app.Listen(":5000")
}
