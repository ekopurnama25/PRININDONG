package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ConfirmasiAdminPesananUsers(c *fiber.Ctx) error {
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
	pesananId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var pesananusers []models.CetakPrintDokumentUsers
	resultsaldo := config.DB.Where("id = ?", pesananId).First(&pesananusers)
	if resultsaldo.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update saldo data",
		})
	}

	// Parse request body to get the data to update saldo
	var requestBodyPesanan struct {
		Lembar           string `json:"lembar"`
		TotalHarga       string `json:"total_harga"`
		Dibayar          string `json:"dibayar"`
		TotalCetak       string `json:"totalcetak"`
		StatusPembayaran string `json:"statuspembayaran"`
		StatusPercetakan string `json:"statuspercetakan"`
	}
	if err := c.BodyParser(&requestBodyPesanan); err != nil {
		return err
	}
	var tinta models.TintaPrint
	if err := config.DB.First(&tinta, pesananusers[0].TintaId).Error; err != nil {
		// Handle the error
		fmt.Println("Error retrieving tinta:", err)
		return err
	}
	TintaCetakanHarga := tinta.HargaPrint
	HargaIdTinta, err := strconv.ParseFloat(TintaCetakanHarga, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing HargaPrint to float:", err)
		return err
	}

	PrintLembar, err := strconv.ParseFloat(requestBodyPesanan.Lembar, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing Lembar to float:", err)
		return err
	}

	BanyakCetak, err := strconv.ParseFloat(requestBodyPesanan.TotalCetak, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing TotalCetak to float:", err)
		return err
	}

	HargaKeseluruhan := HargaIdTinta * PrintLembar * BanyakCetak
	TotalHargaString := strconv.FormatFloat(HargaKeseluruhan, 'f', -1, 64)
	log.Println(TotalHargaString)

	data := models.CetakPrintDokumentUsers{
		Lembar:           requestBodyPesanan.Lembar,
		TotalHarga:       TotalHargaString,
		Dibayar:          "0",
		TotalCetak:       requestBodyPesanan.TotalCetak,
		StatusPembayaran: "Belum Lunas",
		StatusPercetakan: "Sedang diproses Admin",
	}

	// Update saldo in database
	result := config.DB.Model(&models.CetakPrintDokumentUsers{}).Where("id = ?", pesananId).Updates(&data)

	if result.Error != nil {
		return result.Error
	}

	updatedData := models.CetakPrintDokumentUsers{}
	config.DB.First(&updatedData, pesananId)

	// Return the updated data in the response
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully updated data",
		"data":    updatedData, // Return the updated data instead of the result
	})
}

func ConfirmasiUsersPesananUsers(c *fiber.Ctx) error {
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
	pesananId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var pesananusers []models.CetakPrintDokumentUsers
	resultsaldo := config.DB.Where("id = ?", pesananId).First(&pesananusers)
	if resultsaldo.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update saldo data",
		})
	}

	// Parse request body to get the data to update saldo
	var requestBodyPesanan struct {
		TotalHarga       string `json:"total_harga"`
		TotalCetak       string `json:"totalcetak"`
		Dibayar          string `json:"dibayar"`
		StatusPembayaran string `json:"statuspembayaran"`
		StatusPercetakan string `json:"statuspercetakan"`
	}
	if err := c.BodyParser(&requestBodyPesanan); err != nil {
		return err
	}
	var tinta models.TintaPrint
	if err := config.DB.First(&tinta, pesananusers[0].TintaId).Error; err != nil {
		// Handle the error
		fmt.Println("Error retrieving tinta:", err)
		return err
	}
	TintaCetakanHarga := tinta.HargaPrint
	HargaIdTinta, err := strconv.ParseFloat(TintaCetakanHarga, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing HargaPrint to float:", err)
		return err
	}

	PrintLembar, err := strconv.ParseFloat(pesananusers[0].Lembar, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing Lembar to float:", err)
		return err
	}
	totalCetakTrimmed := strings.TrimSpace(requestBodyPesanan.TotalCetak)

	// Validate TotalCetak field
	if totalCetakTrimmed == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "TotalCetak field is empty",
		})
	}

	// Check if TotalCetak contains only numeric characters
	if match, _ := regexp.MatchString(`^\d*\.?\d+$`, totalCetakTrimmed); !match {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "TotalCetak field contains invalid characters",
		})
	}

	// Parse TotalCetak into float
	BanyakCetak, err := strconv.ParseFloat(totalCetakTrimmed, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error parsing TotalCetak to float:", err)
		return err
	}

	HargaKeseluruhan := HargaIdTinta * PrintLembar * BanyakCetak
	TotalHargaString := strconv.FormatFloat(HargaKeseluruhan, 'f', -1, 64)
	log.Println(TotalHargaString)

	var saldo []models.SaldoUsers
	jumlahsaldo := config.DB.Where("user_id  = ?", users.Id).First(&saldo)
	if jumlahsaldo.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": resultsaldo.Error.Error(),
		})
	}
	Saldo := saldo[0].SaldoUser
	sisaSaldoNow, err := strconv.ParseFloat(Saldo, 64)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		return err
	}

	data := models.CetakPrintDokumentUsers{
		TotalHarga:       TotalHargaString,
		Dibayar:          TotalHargaString,
		TotalCetak:       requestBodyPesanan.TotalCetak,
		StatusPembayaran: "Lunas",
		StatusPercetakan: "Sedang Proses Cetak",
	}

	TotHarga, err := strconv.ParseFloat(data.TotalHarga, 64)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		return err
	}

	log.Println(sisaSaldoNow < TotHarga)

	if sisaSaldoNow < TotHarga {
		return c.JSON(fiber.Map{
			"message": "saldo anda tidak cukup",
		})
	} else {

		result := config.DB.Model(&models.CetakPrintDokumentUsers{}).Where("id = ?", pesananId).Updates(&data)

		if result.Error != nil {
			return c.JSON(fiber.Map{
				"message": "Gagal Menambah Data",
			})
		}

		if resultsaldo.RowsAffected < 1 {
			return c.JSON(fiber.Map{
				"message": "Gagal Menambah Data",
			})
		} else {
			sisaSaldoStr := saldo[0].SaldoUser
			sisaSaldo, err := strconv.ParseFloat(sisaSaldoStr, 64)
			if err != nil {
				// handle error
			}

			jumlahStr := data.Dibayar
			jumlah, err := strconv.ParseFloat(jumlahStr, 64)
			if err != nil {
				// handle error
			}

			remainingBalance := sisaSaldo - jumlah
			log.Println(remainingBalance)
			updateSaldo := models.SaldoUsers{
				SaldoUser: strconv.FormatFloat(remainingBalance, 'f', -1, 64),
				UserId:    users.Id,
			}
			config.DB.Model(&saldo).Updates(updateSaldo)
		}
	}

	// Update saldo in database
	result := config.DB.Model(&models.CetakPrintDokumentUsers{}).Where("id = ?", pesananId).Updates(&data)

	if result.Error != nil {
		return result.Error
	}

	updatedData := models.CetakPrintDokumentUsers{}
	config.DB.First(&updatedData, pesananId)

	// Return the updated data in the response
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully updated data",
		"data":    updatedData, // Return the updated data instead of the result
	})
}

func ConfirmasiStatusSelesai(c *fiber.Ctx) error {
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
	pesananId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var pesananusers []models.CetakPrintDokumentUsers
	resultsaldo := config.DB.Where("id = ?", pesananId).First(&pesananusers)
	if resultsaldo.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "failed update saldo data",
		})
	}

	// Parse request body to get the data to update saldo
	var requestBodyPesanan struct {
		StatusPercetakan string `json:"statuspercetakan"`
	}
	if err := c.BodyParser(&requestBodyPesanan); err != nil {
		return err
	}

	data := models.CetakPrintDokumentUsers{
		StatusPercetakan: "Selesai, silahkan ambil pesanan keadmin",
	}

	// Update saldo in database
	result := config.DB.Model(&models.CetakPrintDokumentUsers{}).Where("id = ?", pesananId).Updates(&data)

	if result.Error != nil {
		return result.Error
	}

	updatedData := models.CetakPrintDokumentUsers{}
	config.DB.First(&updatedData, pesananId)

	// Return the updated data in the response
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "successfully updated data",
		"data":    updatedData, // Return the updated data instead of the result
	})
}
