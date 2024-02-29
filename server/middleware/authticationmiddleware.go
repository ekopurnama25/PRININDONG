package middlewares

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsUserAuthenticated(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	token := ctx.Cookies(utils.CookieName)

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else if token != "" {
		tokenString = ctx.Cookies(utils.CookieName)
	}

	if tokenString == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	Token, err := utils.ParseToken(tokenString)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	var user models.Users
	config.DB.Where("id = ?", Token).First(&user)
	return ctx.Next()
}
