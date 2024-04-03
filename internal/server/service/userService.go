package service

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct {
	usecase domain.IUserUseCase
}

func NewUserService(usecase domain.IUserUseCase) *UserService {
	return &UserService{usecase: usecase}
}

func (u *UserService) AuthLoginHandler(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := u.usecase.AuthLogin(&user)
	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
