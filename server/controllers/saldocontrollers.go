package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AllSaldo(c *fiber.Ctx) error {
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
	var saldo []models.SaldoUsers
	config.DB.Preload("Users").Find(&saldo)
	return c.JSON(saldo)
}

func SaveSaldoUsers(c *fiber.Ctx) error {
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
	InputSaldo := struct {
		SaldoUser string `json:"saldo_users"`
		UserId    uint   `json:"usersId"`
	}{}

	if err := c.BodyParser(&InputSaldo); err != nil {
		return err
	}
	createSaldo := models.SaldoUsers{
		SaldoUser: InputSaldo.SaldoUser,
		UserId:    InputSaldo.UserId,
	}
	if err := c.BodyParser(&createSaldo); err != nil {
		return err
	}
	result := config.DB.Create(&createSaldo)

	if result.Error != nil {
		return result.Error
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully added saldo data",
		"data":    createSaldo,
	})
}

func GetIdSaldo(c *fiber.Ctx) error {
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
	saldoId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var saldo models.SaldoUsers
	err = config.DB.Preload("Users").First(&saldo, saldoId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed displays user data",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully get saldo data",
		"data":    saldo,
	})
}

func DeleteIdSaldo(c *fiber.Ctx) error {
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

	saldoId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	result := config.DB.Delete(&models.SaldoUsers{}, saldoId)
	if result.Error != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed delete saldo data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully delete saldo",
		"id":      saldoId,
	})
}

func UpdateSaldoUsers(c *fiber.Ctx) error {
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
	saldoId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var saldo models.SaldoUsers
	checkId := config.DB.First(&saldo, saldoId).Error
	if checkId != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update saldo data",
		})
	}

	if err := c.BodyParser(&saldo); err != nil {
		return err
	}
	saldo.Id = uint(saldoId)
	result := config.DB.Where("id = ?", saldoId).Updates(&saldo)
	if result.Error != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully update saldo data",
		"data":    saldo,
	})
}

func IsiSaldoUsers(c *fiber.Ctx) error {
	cookie := c.Cookies(utils.CookieName)
	if strings.HasPrefix(cookie, "Bearer ") {
		cookie = strings.TrimPrefix(cookie, "Bearer ")
	}
	usersId, err := utils.ParseToken(cookie)
	if err != nil {
		return err
	}
	var users models.Users
	err = config.DB.Preload("Role").First(&users, usersId).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "cookies yang diterima tidak benar",
		})
	}
	userId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var saldo []models.SaldoUsers
	resultsaldo := config.DB.Where("user_id = ?", userId).First(&saldo)
	if resultsaldo.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update saldo data",
		})
	}

	// Parse request body to get the data to update saldo
	var requestBody struct {
		SaldoUser string `json:"saldo_users"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	//Validate that SaldoUser is not empty
	if requestBody.SaldoUser == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "SaldoUser tidak boleh kosong",
		})
	}

	SaldoAwal := strings.TrimSpace(saldo[0].SaldoUser)
	saldoFloat, err := strconv.ParseFloat(SaldoAwal, 64)
	if err != nil {
		return err
	}
	// Convert saldo user strings to integers
	saldoAkhir := strings.TrimSpace(requestBody.SaldoUser)
	requestedSaldo, err := strconv.ParseFloat(saldoAkhir, 64)
	if err != nil {
		return err
	}

	// Calculate new saldo
	newSaldo := saldoFloat + requestedSaldo

	log.Println(newSaldo)
	// Update saldo in database
	// result := config.DB.Model(&saldo).Where("user_id = ?", userId).Update("saldo_user", newSaldo)
	// if result.Error != nil {
	// 	return result.Error
	// }

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully update saldo data",
		"data":    newSaldo,
	})
}
