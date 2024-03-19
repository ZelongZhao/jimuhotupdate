package middlewares

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"strings"
	"time"
)

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
		mc, err := auth.ParseJwtToken(authHeader)
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