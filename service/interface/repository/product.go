package repository

import "format-clean-architecture/entity"

type IProductRepo interface {
	CreateProduct(*entity.Product) (*entity.Product, error)
	GetProduct(id int64) (*entity.Product, error)
	DeleteProduct(id int64) (error)
}