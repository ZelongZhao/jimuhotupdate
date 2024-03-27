package services

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	usecase domain.IUserUseCase
}

func NewUserService(usecase domain.IUserUseCase) *User {
	return &User{usecase: usecase}
}

func (u *User) AuthLoginHandler(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := u.usecase.AuthLogin(&user)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}
