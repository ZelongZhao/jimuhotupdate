package repo

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.IUserRepo {
	return &user{db: db}
}

func (u *user) WithByUsername(username string) domain.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	}
}

func (u *user) GetUser(opts ...domain.DBOption) (*domain.User, error) {
	var user domain.User
	db := u.db
	for _, opt := range opts {
		db = opt(db)
	}
	return &user, db.First(&user).Error
}

func (u *user) GetUserByUsername(username string) (*domain.User, error) {
	return u.GetUser(u.WithByUsername(username))
}

func (u *user) CreateUser(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *user) UpdateUser(user *domain.User) error {
	return u.db.Save(user).Error
}

func (u *user) DelUser(user string) error {
	return u.db.Delete(&domain.User{Username: user}).Error
}
