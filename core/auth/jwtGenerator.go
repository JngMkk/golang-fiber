package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/JngMkk/golang-fiber/core/cache"
	"github.com/JngMkk/golang-fiber/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Generate JWT Tokens(access, refresh)
func GenereateTokens(id uint) (string, string, error) {
	accessToken, err := generateAccessToken(id)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
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

func SetRefreshTokenCookie(c *fiber.Ctx, token string) *fiber.Ctx {
	cookie := new(fiber.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 1)
	cookie.HTTPOnly = true // JavaScript 접근 방지
	// cookie.Secure = true   // HTTPS를 통해서만 전송
	cookie.SameSite = "Lax"
	c.Cookie(cookie)

	return c
}
