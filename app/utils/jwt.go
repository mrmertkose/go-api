package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

//var exprationTime = os.Getenv("JWT_SECRET_KEY_EXPIRE")

func GenerateToken(id uint) (string, error) {
	//expires time sonra bak
	claims := jwt.MapClaims{
		"Id":        id,
		"ExpiresAt": time.Now().Add(1500 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
