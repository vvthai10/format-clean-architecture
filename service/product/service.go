package product

import (
	"format-clean-architecture/entity"
	"format-clean-architecture/service/interface/repository"
	"format-clean-architecture/service/interface/usecase"
)

type Service struct {
	repo repository.IProductRepo
}

func NewService(r repository.IProductRepo) usecase.IProductService{
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateProduct(name, email, password string) (*entity.Product, error){
	return s.repo.CreateProduct(&entity.Product{})
}

func (s *Service) GetProduct(id int64) (*entity.Product, error){
	return s.repo.GetProduct(id)
}

func (s *Service) DeleteProduct(id int64) (error){
	return s.repo.DeleteProduct(id)
}