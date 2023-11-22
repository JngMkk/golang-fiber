package auth

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/JngMkk/golang-fiber/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

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

// Generate JWT Tokens(access, refresh)
func GenereateTokens(id uint) (*Tokens, error) {
	accessToken, err := generateToken(id, time.Minute*time.Duration(config.AccessTokenExpire))
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(id, time.Hour*time.Duration(config.RefreshTokenExpire))
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// Generate Token
func generateToken(id uint, expireTime time.Duration) (string, error) {
	// set claims
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

type TokenData struct {
	UserID  int
	Expires int
}

// Get data by token
func GetTokenData(c *fiber.Ctx) (*TokenData, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := int(claims["sub"].(float64))
		expires := int(claims["exp"].(float64))

		return &TokenData{
			UserID:  userID,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	token := getToken(c)

	t, err := jwt.Parse(token, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getToken(c *fiber.Ctx) string {
	auth := c.Get("Authorization")

	token := strings.Split(auth, " ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.JWTSecret), nil
}
