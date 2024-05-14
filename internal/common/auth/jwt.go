package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	//secret = os.Getenv("JWT_SECRET")
	secret = []byte("privet")
)

type Claims struct {
	ID   int64  `json:"id"`
	Type int16  `json:"type"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWT(claims Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format: expected *Claims, got %T", token.Claims)
	}

	return claims, nil
}
