package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	//secret = os.Getenv("JWT_SECRET")
	secret = []byte("privet")
)

type Claims struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GenerateJWT(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   claims.ID,
		"name": claims.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	return token.SignedString(secret)
}

func ValidateJWT(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(int64)
		name := claims["id"].(string)
		return Claims{
			ID:   id,
			Name: name,
		}, nil
	} else {
		return Claims{}, err
	}
}
