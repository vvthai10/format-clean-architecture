package user

import (
	"fmt"
	"format-clean-architecture/entity"
	"format-clean-architecture/service/interface/repository"
	"format-clean-architecture/service/interface/usecase"
)

type Service struct {
	repo repository.IUserRepo
}

func NewService(r repository.IUserRepo) usecase.IUserService{
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(name, email, password string) (*entity.User, error){
	fmt.Println("SERVICE: Create user")
	return s.repo.CreateUser(&entity.User{})
}

func (s *Service) GetUser(id int64) (*entity.User, error){
	fmt.Println("SERVICE: Get user")
	return s.repo.GetUser(id)
}

func (s *Service) DeleteUser(id int64) (error){
	return s.repo.DeleteUser(id)
}