package usecase

import (
	"errors"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type userUsecase struct {
	userRepo domain.IUserRepo
}

func NewUserUseCase(repo domain.IUserRepo) domain.IUserUseCase {
	return &userUsecase{userRepo: repo}
}

const TokenExpireDuration = time.Hour * 24

func genJwtToken(username string) (string, error) {
	claims := domain.LoginClaims{}
	claims.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	claims.StandardClaims.Issuer = "leoric"
	claims.Username = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(domain.JwtSecret)
}

func (u *userUsecase) AuthLogin(user *domain.User) (string, error) {
	userDB, err := u.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if userDB.Password != user.Password {
		return "", errors.New("wrong password")
	}

	tokenString, err := genJwtToken(user.Username)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *userUsecase) CreateUser(user *domain.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) UpdateUser(user *domain.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *userUsecase) DelUser(user string) error {
	return u.userRepo.DelUser(user)
}

func (u *userUsecase) GetUser(username string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}
