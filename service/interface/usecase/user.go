package usecase

import "format-clean-architecture/entity"

type IUserService interface {
	CreateUser(name, email, password string) (*entity.User, error)
	GetUser(id int64)(*entity.User, error)
	DeleteUser(id int64)(error)
}