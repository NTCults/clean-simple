package repos

import "clean/models"

type UserRepository interface {
	CreateUser(name string) (int, error)
	GetUser(ID int) models.User
}
