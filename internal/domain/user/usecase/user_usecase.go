package usecase

import (
	"cronos/internal/domain/user/entity"
	"cronos/internal/domain/user/repository"
	"errors"
)

type UserUsecase struct {
	UserRepository *repository.UserRepositoryImpl
}

func (uc *UserUsecase) Execute(email, password string) (*entity.User, error) {
	existsEmail, err := uc.UserRepository.FindByEmail(email)
	if existsEmail != nil {
		return &entity.User{}, errors.New("Email already exists")
	}

	newUser, err := entity.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	err = uc.UserRepository.Save(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
