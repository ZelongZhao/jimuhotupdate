package v1

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthLoginHandler(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.FindProfileByUsername(user.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.GenJwtToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
