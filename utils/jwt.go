package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const jwtKey = "changeMe"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(jwtKey))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}
	//email, _ := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
