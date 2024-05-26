package service

import (
	"context"
	v1 "git.vfeda.com/vfedabackendteam/jimuhotupdate/api/hotUpdate/v1"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/domain"
)

type UserService struct {
	usecase domain.IUserUseCase
}

func NewUserService(usecase domain.IUserUseCase) *UserService {
	return &UserService{usecase: usecase}
}

//func (u *UserService) AuthLoginHandler(c *gin.Context) {
//	user := domain.User{}
//	if err := c.ShouldBind(&user); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	tokenString, err := u.usecase.AuthLogin(&user)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"token": tokenString})
//}

func (u *UserService) Login(ctx context.Context, loginRequest *v1.LoginRequest) (*v1.LoginResponse, error) {
	user := domain.User{}
	user.Username = loginRequest.Username
	user.Password = loginRequest.Password
	tokenString, err := u.usecase.AuthLogin(&user)
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse{Token: tokenString}, nil
}
