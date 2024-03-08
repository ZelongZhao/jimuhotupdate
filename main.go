package main

import (
	"errors"
	"git.vfeda.com/vfeda/JiMuHotUpdate/Middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Profile struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getProfile(user string) (Profile, error) {
	if user == "leoric" {
		return Profile{
			Username: "leoric",
			Password: "123",
		}, nil
	}
	return Profile{}, errors.New("user not found")
}

func AuthLoginHandler(c *gin.Context) {
	var user Profile
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if _, err := getProfile(user.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tokenString, err := Middlewares.GenJwtToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func main() {
	// 1.创建路由
	r := gin.Default()

	r.Run(":8080")
}
