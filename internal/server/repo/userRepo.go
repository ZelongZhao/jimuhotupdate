package repo

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.IUserRepo {
	return &userRepo{db: db}
}

func (u *userRepo) WithByUsername(username string) domain.DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	}
}

func (u *userRepo) GetUser(opts ...domain.DBOption) (*domain.User, error) {
	var user domain.User
	db := u.db
	for _, opt := range opts {
		db = opt(db)
	}
	return &user, db.First(&user).Error
}

func (u *userRepo) GetUserByUsername(username string) (*domain.User, error) {
	return u.GetUser(u.WithByUsername(username))
}

func (u *userRepo) CreateUser(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) UpdateUser(user *domain.User) error {
	return u.db.Save(user).Error
}

func (u *userRepo) DelUser(user string) error {
	return u.db.Delete(&domain.User{Username: user}).Error
}
