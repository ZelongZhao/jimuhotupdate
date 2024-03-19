package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	JwtSecret = []byte("secret")
)

const TokenExpireDuration = time.Hour * 24

type TestClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenJwtToken(username string) (string, error) {
	claims := TestClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "leoric",
		},
		username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseJwtToken(jwtToken string) (*TestClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &TestClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*TestClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims type")
	}
}
