package auth

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/JngMkk/golang-fiber/core/cache"
	"github.com/JngMkk/golang-fiber/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Generate JWT Tokens(access, refresh)
func GenereateTokens(id uint) (*Tokens, error) {
	accessToken, err := generateAccessToken(id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken(id)
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// Generate Access Token
func generateAccessToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Minute * time.Duration(config.AccessTokenExpire)).Unix(),
	}
	token, err := generateTokenString(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Generate Refresh Token
func generateRefreshToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * time.Duration(config.RefreshTokenExpire)).Unix(),
	}
	token, err := generateTokenString(claims)
	if err != nil {
		return "", err
	}
	if err := saveRefreshToken(id, token); err != nil {
		return "", err
	}

	return token, nil
}

func generateTokenString(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

// Save Refresh Token in Redis
func saveRefreshToken(id uint, token string) error {
	ctx := context.Background()
	rConn := cache.Connect()
	err := rConn.Set(ctx, strconv.FormatUint(uint64(id), 10), token, time.Hour*1).Err()
	return err
}

type TokenData struct {
	UserID  uint
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
		userID := uint(claims["sub"].(float64))
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
