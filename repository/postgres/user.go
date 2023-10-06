package postgres

import (
	"fmt"
	"format-clean-architecture/entity"
	"format-clean-architecture/service/interface/repository"

	"gorm.io/gorm"
)

type UserRepo struct{
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.IUserRepo{
	return &UserRepo{
		db: db,
	}
}

func (q *UserRepo) CreateUser(u *entity.User) (*entity.User, error){
	fmt.Println("REPO: Create user")
	return &entity.User{}, nil
}

func (q *UserRepo) GetUser(id int64) (*entity.User, error){
	fmt.Println("REPO: Get user")
	return &entity.User{}, nil
}

func (q *UserRepo) DeleteUser(id int64) (error){
	return nil
}