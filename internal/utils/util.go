package utils

import (
	"errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(errors.New("unable hash password"))
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
