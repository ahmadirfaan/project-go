package utils

import (
	"errors"
	"github.com/ahmadirfaan/project-go/app"
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

// This function to generates Access Token
func GenerateToken(user database.User) (*string, string, error) {
	app := app.Init()
	jwtSecretKey := app.Config.JWTSecret

	// Generates Access Token Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["userId"] = user.Id
	claims["exp"] = time.Now().UTC().Add(time.Hour * 24 * 7).Unix()
	time := time.Unix(time.Now().UTC().Add(time.Hour*24*7).Unix(), 0).Format(time.RFC3339)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return nil, time, err
	}
	return &t, time, err
}

func ExtractToken(c *fiber.Ctx) (string, error) {
	appFiber := app.Init()
	tokenString := c.Get("Authorization")
	tokenString = strings.ReplaceAll(tokenString, " ", "")
	tokenString = strings.ReplaceAll(tokenString, "Bearer", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(appFiber.Config.JWTSecret), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)
	return strconv.Itoa(int(userId)), err
}
