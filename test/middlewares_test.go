package main

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"

	"git.vfeda.com/vfeda/JiMuHotUpdate/Middlewares"
)

func TestGenJwtToken(t *testing.T) {
	username := "testuser"
	expectedExpire := time.Now().Add(Middlewares.TokenExpireDuration).Unix()
	expectedIssuer := "leoric"
	expectedUsername := "testuser"

	token, err := Middlewares.GenJwtToken(username)
	if err != nil {
		t.Errorf("Failed to generate JWT token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &Middlewares.TestClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Middlewares.JwtSecret, nil
	})
	if err != nil {
		t.Errorf("Failed to parse JWT token: %v", err)
	}

	claims := parsedToken.Claims.(*Middlewares.TestClaims)

	if claims.ExpiresAt != expectedExpire {
		t.Errorf("Expected expiration time %v, but got %v", expectedExpire, claims.ExpiresAt)
	}
	if claims.Issuer != expectedIssuer {
		t.Errorf("Expected issuer %v, but got %v", expectedIssuer, claims.Issuer)
	}
	if claims.Username != expectedUsername {
		t.Errorf("Expected username %v, but got %v", expectedUsername, claims.Username)
	}
}

func TestParseJwtToken_ValidToken(t *testing.T) {
	//jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMTA0MjAsImp0aSI6IjEyMzQ1Njc4OTAiLCJpYXQiOjE3MTAyMTMyMjAsImlzcyI6Imlzc3VlciIsInVzZXJuYW1lIjoiIn0.wUlkNvbaLFexnYzyekDx7A2GE4OKlj_tmEmbn1f18Pk"

	expectedClaims := &Middlewares.TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(Middlewares.TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Add(-time.Hour * 3).Unix(),
			Id:        "1234567890",
			Issuer:    "issuer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expectedClaims)
	jwtToken, _ := token.SignedString(Middlewares.JwtSecret)

	claims, err := Middlewares.ParseJwtToken(jwtToken)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if claims == nil {
		t.Error("Expected claims not to be nil, but it was")
	}

	if claims.ExpiresAt != expectedClaims.ExpiresAt {
		t.Errorf("Expected expires at %d, but got %d", expectedClaims.ExpiresAt, claims.ExpiresAt)
	}

	if claims.IssuedAt != expectedClaims.IssuedAt {
		t.Errorf("Expected issued at %d, but got %d", expectedClaims.IssuedAt, claims.IssuedAt)
	}

	if claims.Id != expectedClaims.Id {
		t.Errorf("Expected jti %s, but got %s", expectedClaims.Id, claims.Id)
	}

	if claims.Issuer != expectedClaims.Issuer {
		t.Errorf("Expected issuer %s, but got %s", expectedClaims.Issuer, claims.Issuer)
	}
}
