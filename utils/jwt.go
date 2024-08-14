package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "assa221@3$%ASASSAsas87"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signed in method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Couldnt parse error")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid Tokens")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil
}
