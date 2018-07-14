package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func MakeToken(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain. In this case we are storing the username
	// and the expiration date.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expires":  time.Now().Add(72 * time.Hour),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("thisismyveysecretkey!!@@33$$5asqweasdqwe"))
	return tokenString, err
}
