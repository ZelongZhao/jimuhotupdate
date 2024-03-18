package v1

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/Middlewares"
	"git.vfeda.com/vfeda/JiMuHotUpdate/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthLoginHandler(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.GetProfile(user.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := Middlewares.GenJwtToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
