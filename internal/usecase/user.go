package usecase

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
)

type user struct {
	repo domain.IUserRepo
}

func NewUserUseCase(repo domain.IUserRepo) domain.IUserUseCase {
	return &user{repo: repo}
}

func (u *user) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *user) UpdateUser(user *domain.User) error {
	return u.repo.UpdateUser(user)
}

func (u *user) DelUser(user string) error {
	return u.repo.DelUser(user)
}

func (u *user) GetUser(username string) (*domain.User, error) {
	return u.repo.GetUserByUsername(username)
}
