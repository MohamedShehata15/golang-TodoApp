package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func HashingPassword(plainPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}