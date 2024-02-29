package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
const signingKey = "SomeSecretKey" // TODO: move it to a safe place
func CreateToken(issuer string, expirationTime time.Time) (string, error) {
	env:= goDotEnvVariable("TOKEN_SCRET")
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env))
}

func CreateRefreshToken(issuer string, expirationTime time.Time) (string, error) {
	env:= goDotEnvVariable("TOKEN_SCRET_REFRESH")
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env))
}

func ParseToken(tokenString string) (string, error) {
	env:= goDotEnvVariable("TOKEN_SCRET")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(env), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims) // Casting the token.Claims to the struct jwt.StandardClaims

	return claims.Issuer, nil
}


func ParseRefreshToken(tokenString string) (string, error) {
	env:= goDotEnvVariable("TOKEN_SCRET_REFRESH")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(env), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims) // Casting the token.Claims to the struct jwt.StandardClaims
	return claims.Issuer, nil
}