package Middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"strings"
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

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		mc, err := ParseJwtToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}
func RateLimitMiddleware() func(c *gin.Context) {
	bucket := ratelimit.NewBucketWithQuantum(time.Minute, 10, 1)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden})
			c.Abort()
			return
		}
		c.Next()
	}
}
