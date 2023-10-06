package usecase

import "format-clean-architecture/entity"

type IProductService interface {
	CreateProduct(name, email, password string) (*entity.Product, error)
	GetProduct(id int64)(*entity.Product, error)
	DeleteProduct(id int64)(error)
}