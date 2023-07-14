package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("cloudservice")

func NewToken() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "CloudService",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("No match algrithem: %s", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else {
		return false, token.Claims.Valid()
	}
}
