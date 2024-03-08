package Middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	jwtSecret = []byte("secret")
)

const TokenExpireDuration = time.Hour * 24

type TestClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func (t *TestClaims) Vaild() {
	return
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
	return token.SignedString(jwtSecret)
}

//func ParseToken(tokenString string) (*MyClaims, error) {
//	// 解析token
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
//		return MySecret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}

func ParseJwtToken(jwtToken string) (*TestClaims, error) {
	//token, err := jwt.ParseWithClaims(jwtToken, &TestClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return jwtSecret, nil
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return nil, err
}
