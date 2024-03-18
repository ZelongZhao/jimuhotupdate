package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetProfile(user string) (User, error) {
	if user == "leoric" {
		return User{
			Username: "leoric",
			Password: "123",
		}, nil
	}
	return User{}, errors.New("user not found")
}
