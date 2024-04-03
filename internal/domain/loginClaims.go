package domain

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	JwtSecret = []byte("secret")
)

type LoginClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}