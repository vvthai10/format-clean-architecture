package repository

import "format-clean-architecture/entity"

type IUserRepo interface {
	CreateUser(*entity.User) (*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	DeleteUser(id int64) (error)
}