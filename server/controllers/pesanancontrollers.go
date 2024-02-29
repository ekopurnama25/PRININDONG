package controllers

import (
	"apk-chat-serve/config"
	"apk-chat-serve/models"
	"apk-chat-serve/utils"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func generateNumericSuffix(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	suffix := make([]byte, length)
	for i := range suffix {
		suffix[i] = charset[rand.Intn(len(charset))]
	}
	return string(suffix)
}

func SavePesanan(c *fiber.Ctx) error {
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

	file, err := c.FormFile("FilenameCetak")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	PesananCetakPrintnin := struct {
		NoPesananan      string `json:"no_pesanan"`
		JudulPesanan     string `json:"judul_pesanan"`
		Lembar           string `json:"lembar"`
		TotalHarga       string `json:"total_harga"`
		Dibayar          string `json:"dibayar"`
		StatusPembayaran string `json:"statuspembayaran"`
		StatusPercetakan string `json:"statuspercetakan"`
		TanggalPemesanan string `json:"tanggalpemesanan"`
		TintaId          uint   `json:"tinta_id"`
	}{}

	if err := c.BodyParser(&PesananCetakPrintnin); err != nil {
		return err
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	parts := strings.Split(file.Filename, ".")
	if len(parts) < 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid file name",
		})
	}
	fileExt := parts[len(parts)-1]

	// Define the allowed file types
	allowedExts := map[string]bool{
		"xls":  true,
		"xlsx": true,
		"pdf":  true,
		"doc":  true,
		"docx": true,
	}

	// Check if the file type is supported
	if !allowedExts[strings.ToLower(fileExt)] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Unsupported file type",
		})
	}

	err = c.SaveFile(file, fmt.Sprintf("./utils/file/%s.%s", filename, fileExt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to save file",
		})
	}

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("./utils/file/%s", image))
	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	imageUrl := fmt.Sprintf("http://localhost:5000/utils/file/%s", image)

	var saldo []models.SaldoUsers
	resultsaldo := config.DB.Where("user_id = ?", users.Id).First(&saldo)
	if resultsaldo.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": resultsaldo.Error.Error(),
		})
	}

	tanggal := time.Now().UTC()
	tangalStr := tanggal.String()
	prefix := "INV"

	// Define the length of the numeric suffix
	suffixLength := 6 // Adjust as needed

	// Generate the numeric suffix
	numericSuffix := generateNumericSuffix(suffixLength)

	// Combine prefix and numeric suffix
	automaticCode := prefix + numericSuffix

	data := models.CetakPrintDokumentUsers{
		NoPesananan:      automaticCode,
		JudulPesanan:     PesananCetakPrintnin.JudulPesanan,
		Lembar:           "0",
		TotalHarga:       "0",
		Dibayar:          "0",
		StatusPembayaran: "Belum Lunas",
		StatusPercetakan: "Sedang diproses Admin",
		TanggalPemesanan: tangalStr,
		TintaId:          PesananCetakPrintnin.TintaId,
		FilenameCetak:    image,
		UrlFile:          imageUrl,
		UserId:           users.Id,
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil Menambah Data",
		"data":    data,
	})
}
