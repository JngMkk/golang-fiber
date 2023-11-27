package auth

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(pw string) error {
	if len(pw) < 8 || len(pw) > 16 {
		return errors.New("password must be between 8 and 16 characters")
	}

	hasLower := regexp.MustCompile(`[a-z]`)
	hasCapital := regexp.MustCompile(`[A-Z]`)
	hasDigit := regexp.MustCompile(`\d`)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`)

	if !hasLower.MatchString(pw) {
		return errors.New("password must contain at least one lower letter")
	}
	if !hasCapital.MatchString(pw) {
		return errors.New("password must contain at least one capital letter")
	}
	if !hasDigit.MatchString(pw) {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial.MatchString(pw) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func HashPassword(pw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
