package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	GenerateToken(userID int) (string, error)
	ValidationToken(encodedToken string) (*jwt.Token, error)
}

type auth struct{}

func NewAuth() *auth {
	return &auth{}
}

var SIGNED_KEY = []byte("jJkk98_")

func (a *auth) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString(SIGNED_KEY)
	if err != nil {
		return signedString, err
	}
	return signedString, nil
}

func (a *auth) ValidationToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SIGNED_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
