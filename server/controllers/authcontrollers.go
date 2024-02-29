package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AUthUsersMiddlaware(c *fiber.Ctx) error {
	loginDto := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&loginDto); err != nil {
		return err
	}

	var users models.Users
	result := config.DB.Where("email = ?", loginDto.Email).First(&users)
	if result.RowsAffected > 0 {
		if err := bcrypt.CompareHashAndPassword(users.Password, []byte(loginDto.Password)); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"message": "Incorrect password !"})
		} else {
			expirationTimeRefresh := time.Now().Add(24 * time.Hour)
			expirationTime := time.Now().Add(30 * time.Second)
			Token, err := utils.CreateToken(strconv.Itoa(int(users.Id)), expirationTime)
			Refresh, err := utils.CreateRefreshToken(strconv.Itoa(int(users.Id)), expirationTimeRefresh)
			if err != nil {
				return err
			}

			var tokens []models.AuthUserTokens
			resultToken := config.DB.Where("user_id = ?", users.Id).First(&tokens)

			if resultToken.RowsAffected < 1 {
				tokenscreate := models.AuthUserTokens{
					AccessToken: Token,
					RefeshToken: Refresh,
					UserId:      users.Id,
				}
				config.DB.Create(&tokenscreate)
			} else {
				tokenscreate := models.AuthUserTokens{
					AccessToken: Token,
					RefeshToken: Refresh,
					UserId:      users.Id,
				}
				config.DB.Model(&tokens).Updates(tokenscreate)
				//database.DB.Updates(&tokenscreate)
			}

			c.Cookie(&fiber.Cookie{
				Name:     utils.CookieName,
				Value:    Token,
				Path:     "",
				Domain:   "",
				MaxAge:   0,
				Expires:  time.Now().Add(30 * time.Second),
				Secure:   false,
				HTTPOnly: true,
				SameSite: "lax",
			})

			return c.Status(200).JSON(fiber.Map{
				"status":       200,
				"message":      "AJG HARDCODE!",
				"AccessToken":  Token,
				"RefreshToken": Refresh,
				"UserId":       users.Id,
			})
		}
	} else {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed login users data",
		})
	}
}

func PostRefreshToken(c *fiber.Ctx) error {
	RefToken := struct {
		RefeshToken string `json:"refeshtoken"`
	}{}

	if err := c.BodyParser(&RefToken); err != nil {
		return err
	}
	RefeshTokenServer, err := utils.ParseRefreshToken(RefToken.RefeshToken)
	if err != nil {
		return err
	}
	var token models.AuthUserTokens
	result := config.DB.Where("user_id = ?", RefeshTokenServer).First(&token)
	expirationTimeRefresh := time.Now().Add(24 * time.Hour)
	expirationTime := time.Now().Add(30 * time.Second)
	Token, err := utils.CreateToken(strconv.Itoa(int(token.UserId)), expirationTime)
	Refresh, err := utils.CreateRefreshToken(strconv.Itoa(int(token.UserId)), expirationTimeRefresh)
	if err != nil {
		return err
	}
	if result.RowsAffected > 0 {
		tokenscreate := models.AuthUserTokens{
			AccessToken: Token,
			RefeshToken: Refresh,
			UserId:      token.UserId,
		}
		config.DB.Model(&token).Updates(tokenscreate)
		return c.Status(200).JSON(fiber.Map{
			"status":       200,
			"message":      "successfully refresh token user data",
			"AccessToken":  Token,
			"RefreshToken": Refresh,
		})
	} else {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed refresh token users data",
		})
	}
}

func GetUsersLogin(c *fiber.Ctx) error {
	cookie := c.Cookies(utils.CookieName)
	if strings.HasPrefix(cookie, "Bearer ") {
		cookie = strings.TrimPrefix(cookie, "Bearer ")
	}
	userId, err := utils.ParseToken(cookie)
	if err != nil {
		return err
	}
	var users models.Users
	err = config.DB.Preload("Role").First(&users, userId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "sukses verify token users data",
		"data":    users,
	})
}

func LogoutAuth(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     utils.CookieName,
		Value:    "",
		Path:     "",
		Domain:   "",
		MaxAge:   0,
		Expires:  time.Now().Add(-(2 * time.Hour)),
		Secure:   false,
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.JSON(fiber.Map{"message": "Berhasil Logout"})
}
