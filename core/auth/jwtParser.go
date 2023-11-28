package auth

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/JngMkk/golang-fiber/core/cache"
	"github.com/JngMkk/golang-fiber/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	UserID  uint
	Expires int
}

// Get data by token
func GetAccessTokenData(c *fiber.Ctx) (*TokenData, error) {
	token, err := verifyAccessToken(c)
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

func verifyAccessToken(c *fiber.Ctx) (*jwt.Token, error) {
	token := getAccessToken(c)

	t, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getAccessToken(c *fiber.Ctx) string {
	auth := c.Get("Authorization")

	token := strings.Split(auth, " ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}

// validate refresh token
func ValidateRefreshToken(c *fiber.Ctx) (uint, error) {
	token, err := verifyRefreshToken(c)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := uint(claims["sub"].(float64))
		storedToken, err := getRefreshTokenFromCache(userID)
		if err != nil {
			return 0, err
		}

		tokenString, err := token.SignedString(config.JWTSecret)
		if err != nil {
			return 0, err
		}

		if tokenString == storedToken {
			return userID, nil
		}
	}

	return 0, nil
}

func verifyRefreshToken(c *fiber.Ctx) (*jwt.Token, error) {
	token, err := getRefreshTokenFromCookie(c)
	if err != nil {
		return nil, err
	}

	t, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// refresh token from cookie
func getRefreshTokenFromCookie(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("refreshToken")
	if cookie == "" {
		return "", errors.New("cookie not found")
	}

	return cookie, nil
}

// refresh token from redis
func getRefreshTokenFromCache(id uint) (string, error) {
	ctx := context.Background()
	rConn := cache.Connect()
	storedToken, err := rConn.Get(ctx, strconv.FormatUint(uint64(id), 10)).Result()
	if err != nil {
		return "", err
	}

	return storedToken, nil
}

func parseToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.JWTSecret), nil
}
