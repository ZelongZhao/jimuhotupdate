package services

import "git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"

type User struct {
	repo domain.IUserUseCase
}

func NewUserService(repo domain.IUserUseCase) *User {
	return &User{repo: repo}
}
