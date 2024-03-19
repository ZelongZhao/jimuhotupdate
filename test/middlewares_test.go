package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGenJwtToken(t *testing.T) {
	username := "testuser"
	expectedExpire := time.Now().Add(middlewares.TokenExpireDuration).Unix()
	expectedIssuer := "leoric"
	expectedUsername := "testuser"

	token, err := middlewares.GenJwtToken(username)
	if err != nil {
		t.Errorf("Failed to generate JWT token: %v", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &middlewares.TestClaims{}, func(token *jwt.Token) (interface{}, error) {
		return middlewares.JwtSecret, nil
	})
	if err != nil {
		t.Errorf("Failed to parse JWT token: %v", err)
	}

	claims := parsedToken.Claims.(*middlewares.TestClaims)

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
	expectedClaims := &middlewares.TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(middlewares.TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Add(-time.Hour * 3).Unix(),
			Id:        "1234567890",
			Issuer:    "issuer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expectedClaims)
	jwtToken, _ := token.SignedString(middlewares.JwtSecret)

	claims, err := middlewares.ParseJwtToken(jwtToken)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if claims == nil {
		t.Error("Expected claims not to be nil, but it was")
		return
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

func TestParseJwtToken_InvalidToken(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJCustomFieldIjoiVmFsaWRJZCJ9.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5d"

	_, err := middlewares.ParseJwtToken(jwtToken)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestJWTAuthMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.JWTAuthMiddleware())

	// Test case 1: No authorization header
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401}", w.Body.String())

	// Test case 2: Invalid authorization header
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Set("authorization", "invalid")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401}", w.Body.String())

	// Test case 3: Valid authorization header
	token, _ := middlewares.GenJwtToken("testuser")
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Set("authorization", token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)
	//assert.Equal(t, gin.H{"username": "testuser"}, w.Body.String())
}

func TestRateLimitMiddleware(t *testing.T) {
	rateLimitMiddleware := middlewares.RateLimitMiddleware()

	// 创建一个gin引擎实例
	router := gin.New()
	router.Use(rateLimitMiddleware)

	// 设置路由处理函数以便测试
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// 测试令牌充足的情况
	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
		}
	}

	// 测试令牌不足的情况
	{
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusForbidden {
			t.Errorf("Expected status code %d but got %d when token is not available", http.StatusForbidden, w.Code)
		}
	}
}
