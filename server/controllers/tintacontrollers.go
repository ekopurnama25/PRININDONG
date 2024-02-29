package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AllTinta(c *fiber.Ctx) error {
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
	var tinta []models.TintaPrint
	config.DB.Find(&tinta)
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully get data tinta",
		"data":    tinta,
	})
}

func SaveTinta(c *fiber.Ctx) error {
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
	var tinta models.TintaPrint
	if err := c.BodyParser(&tinta); err != nil {
		return err
	}
	result := config.DB.Create(&tinta)
	if result.Error != nil {
		return result.Error
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully added tinta data",
		"data":    tinta,
	})
}

func GetIdTinta(c *fiber.Ctx) error {
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
	TintaId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var tinta models.TintaPrint
	err = config.DB.First(&tinta, TintaId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed displays tinta data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully displays tinta data",
		"data":    tinta,
	})
}

func DeleteIdTinta(c *fiber.Ctx) error {
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

	tintaId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	result := config.DB.Delete(&models.TintaPrint{}, tintaId)
	if result.Error != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed delete tinta data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully delete tinta data",
		"id":      tintaId,
	})
}

func UpdateTinta(c *fiber.Ctx) error {
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
	tintaId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var tinta models.TintaPrint
	checkId := config.DB.First(&tinta, tintaId).Error
	if checkId != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update tinta data",
		})
	}

	if err := c.BodyParser(&tinta); err != nil {
		return err
	}
	tinta.Id = uint(tintaId)
	result := config.DB.Where("id = ?", tintaId).Updates(&tinta)
	if result.Error != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully update tinta data",
		"data":    tinta,
	})
}
