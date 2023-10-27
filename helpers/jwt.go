package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"id": userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
		"role": "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenAdmin(AdminID uint) (string, error) {
    jwtSecret := []byte(os.Getenv("JWT_SECRET"))

    claims := jwt.MapClaims{
        "id":   AdminID,
        "exp":   time.Now().Add(time.Hour * 1).Unix(),
        "iat":   time.Now().Unix(),
		"role": "admin",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		UserId := claims["id"].(float64)
		return UserId
	}
	return 0
}