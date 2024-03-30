package usecase

import (
	"cronos/internal/domain/user/repository"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserAuthenticateUsecase struct {
	UserRepository *repository.UserRepositoryImpl
}

func (uc *UserAuthenticateUsecase) Execute(email, password string) (string, error) {
	u, err := uc.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("Erro ao autenticar usuário, verifique seu email")
	}

	if !u.ValidatePassword(password) {
		return "", errors.New("Senha incorreta")
	}

	claims := jwt.MapClaims{
		"sub":   u.ID.String(),
		"exp":   time.Now().Add(time.Second * 3000).Unix(),
		"email": u.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
