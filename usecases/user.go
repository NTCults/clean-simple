package usecases

import (
	"clean/models"
	"clean/repos"
)

type UserUsecase interface {
	CreateUser(name string) (int, error)
	GetUser(ID int) models.User
}

type userUsecase struct {
	repo repos.UserRepository
}

func NewUserUsecase(repo repos.UserRepository) *userUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(name string) (int, error) {
	id, err := u.repo.CreateUser(name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userUsecase) GetUser(ID int) models.User {
	user := u.repo.GetUser(ID)
	return user
}
