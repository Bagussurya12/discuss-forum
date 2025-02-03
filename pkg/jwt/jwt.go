package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(10 * time.Minute).Unix(),
		},
	)

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
