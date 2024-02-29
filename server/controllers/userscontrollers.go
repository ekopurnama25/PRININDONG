package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.Users
	config.DB.Preload("Role").Find(&users)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "successfully displays user data",
		"data":    users,
	})
}

func SaveUsers(c *fiber.Ctx) error {
	CreateUsers := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		IdRole   uint   `json:"roleId"`
	}{}
	if err := c.BodyParser(&CreateUsers); err != nil {
		return err
	}
	users := models.Users{
		Username: CreateUsers.Username,
		Email:    CreateUsers.Email,
		IdRole:   CreateUsers.IdRole,
	}

	users.SetPassword(string(CreateUsers.Password))
	result := config.DB.Create(&users)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed to add user data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully added user data",
		"data":    users,
	})
}

func GetIdUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var users models.Users
	err = config.DB.Preload("Role").First(&users, userId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed displays user data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully displays user data",
		"data":    users,
	})
}

func DeleteIdUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	result := config.DB.Delete(&models.Users{}, userId)
	if result.Error != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed delete user data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully delete users data",
		"id":      userId,
	})
}

func UpdateUsers(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var users models.Users
	checkId := config.DB.Preload("Role").First(&users, userId).Error
	if checkId != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update user data",
		})
	}
	if err := c.BodyParser(&users); err != nil {
		return err
	}
	users.Id = uint(userId)
	result := config.DB.Model(&users).Updates(users).Error
	if result != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully update user data",
		"data":    users,
	})
}
