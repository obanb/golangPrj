package service

import (
	"awesomeProject/common"
	"awesomeProject/domain"
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	SignUpUser(request dto.SignUpUserRequest) (*domain.User, *errs.AppError)
}

func (s DefaultUserService) SignUpUser(req dto.SignUpUserRequest) (*domain.User, *errs.AppError) {
	user := domain.User{
		Name:          req.Name,
		Surname:       req.Surname,
		Nickname:      req.Nickname,
		Position:      req.Position,
		AccountStatus: common.Inactive,
		AccountType:   common.Common,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	hash, hashError := GetHash([]byte(req.Password))
	if hashError != nil {
		return nil, hashError
	}

	user.Password = *hash

	newUser, saveError := s.repo.Save(user)

	if saveError != nil {
		return nil, saveError
	}

	return newUser, nil
}

func GetHash(p []byte) (*string, *errs.AppError) {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from")
	}
	stringed := string(hash)
	return &stringed, nil
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
