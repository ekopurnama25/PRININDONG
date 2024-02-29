package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(80);unique" form:"username" binding:"required"`
	Email    string `gorm:"type:varchar(80);unique" form:"email" binding:"required"`
	Password []byte `gorm:"not null" json:"-"`
	IdRole   uint   `gorm:"not null" json:"roleId"`
	Role     Role   `gorm:"foreignKey:IdRole"`
}

type Role struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Role string `gorm:"type:varchar(80)"`
}

type AuthUserTokens struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	AccessToken string `gorm:"type:varchar(350);" form:"access_token" binding:"required"`
	RefeshToken string `gorm:"type:varchar(350);" form:"refesh_token" binding:"required"`
	UserId      uint   `gorm:"not null" json:"usersId"`
	Users       Users  `gorm:"foreignKey:UserId"`
}

type SaldoUsers struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	SaldoUser string `gorm:"type:varchar(350);" form:"saldo_users" binding:"required"`
	UserId    uint   `gorm:"not null" json:"usersId"`
	Users     Users  `gorm:"foreignKey:UserId"`
}

type CetakPrintDokumentUsers struct {
	Id               uint       `gorm:"primaryKey" json:"id"`
	NoPesananan      string     `gorm:"type:varchar(350);" form:"no_pesanan" binding:"required"`
	JudulPesanan     string     `gorm:"type:varchar(350);" form:"judul_pesanan" binding:"required"`
	Lembar           string     `gorm:"type:varchar(350);" form:"lembar" binding:"required"`
	TotalHarga       string     `gorm:"type:varchar(350);" form:"totalharga" binding:"required"`
	Dibayar          string     `gorm:"type:varchar(350);" form:"dibayar" binding:"required"`
	TotalCetak       string     `gorm:"type:varchar(350);" form:"totalcetak" binding:"required"`
	TanggalPemesanan string     `gorm:"type:varchar(350);" form:"tanggalpemesanan" binding:"required"`
	StatusPembayaran string     `gorm:"type:varchar(350);" form:"statuspembayaran" binding:"required"`
	StatusPercetakan string     `gorm:"type:varchar(350);" form:"statuspercetakan" binding:"required"`
	FilenameCetak    string     `gorm:"type:varchar(350);" form:"filename_cetak" binding:"required"`
	UrlFile          string     `gorm:"type:varchar(350);" form:"urlfile" binding:"required"`
	TintaId          uint       `gorm:"not null" json:"tintaId"`
	UserId           uint       `gorm:"not null" json:"usersId"`
	Users            Users      `gorm:"foreignKey:UserId"`
	TintaPrint       TintaPrint `gorm:"foreignKey:TintaId"`
}

type TintaPrint struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	ColorTinta string `gorm:"type:varchar(80);"`
	HargaPrint string `gorm:"type:varchar(80);"`
}

func (users *Users) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	users.Password = hashedPassword
	return nil
}

func (users *Users) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(users.Password, []byte(password))
}
