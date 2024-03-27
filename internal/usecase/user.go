package usecase

import (
	"errors"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type user struct {
	repo domain.IUserRepo
}

func NewUserUseCase(repo domain.IUserRepo) domain.IUserUseCase {
	return &user{repo: repo}
}

var (
	JwtSecret = []byte("secret")
)

const TokenExpireDuration = time.Hour * 24

type TestClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func genJwtToken(username string) (string, error) {
	claims := TestClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "leoric",
		},
		username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func (u *user) AuthLogin(user *domain.User) (string, error) {
	userDB, err := u.repo.GetUserByUsername(user.Username)
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
