package config

import (
	"apk-chat-serve/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/api-print?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.AuthUserTokens{}, &models.Role{}, &models.Users{}, &models.SaldoUsers{}, &models.TintaPrint{}, &models.CetakPrintDokumentUsers{})
	DB = db
}