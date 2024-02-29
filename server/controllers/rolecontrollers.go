package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	var roles []models.Role
	config.DB.Find(&roles)
	return c.JSON(roles)
}

func SaveRoles(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	result := config.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully added roles data",
		"data":    role,
	})
}

func GetIdRoles(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	RoleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var role models.Role
	err = config.DB.First(&role, RoleId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed displays roles data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully displays roles data",
		"data":    role,
	})
}

func DeleteIdRole(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	roleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	result := config.DB.Delete(&models.Role{}, roleId)
	if result.Error != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed delete roles data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully delete roles data",
		"id":      roleId,
	})
}

func UpdateRoles(c *fiber.Ctx) error {
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
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	roleId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var roles models.Role
	checkId := config.DB.First(&roles, roleId).Error
	if checkId != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update roles data",
		})
	}
	if err := c.BodyParser(&roles); err != nil {
		return err
	}
	roles.Id = uint(roleId)
	result := config.DB.Model(&roles).Updates(roles)
	if result.Error != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully update roles data",
		"data":    roles,
	})
}
