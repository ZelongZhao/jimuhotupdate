package domain

import (
	"gorm.io/gorm"
)

//var db = data.JimuDb

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBOption func(*gorm.DB) *gorm.DB

type IUserUseCase interface {
	GetUser(username string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DelUser(user string) error
	AuthLogin(user *User) (string, error)
}

type IUserRepo interface {
	WithByUsername(username string) DBOption
	GetUser(opts ...DBOption) (*User, error)
	GetUserByUsername(username string) (*User, error)

	CreateUser(user *User) error
	UpdateUser(user *User) error
	DelUser(user string) error
}
