package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vmkevv/form_backend/structs"
)

// GenToken generates a token, based in id and email
func GenToken(id int, email string) (string, error) {
	expirationTime := time.Now().Add(131400 * time.Hour)
	claims := &structs.Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(structs.JwtKey)
	return tokenString, err
}
