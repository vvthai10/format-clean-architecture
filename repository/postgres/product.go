package postgres

import (
	"format-clean-architecture/entity"
	"format-clean-architecture/service/interface/repository"

	"gorm.io/gorm"
)

type ProductRepo struct{
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repository.IProductRepo{
	return &ProductRepo{
		db: db,
	}
}

func (q *ProductRepo) CreateProduct(u *entity.Product) (*entity.Product, error){

	return &entity.Product{}, nil
}

func (q *ProductRepo) GetProduct(id int64) (*entity.Product, error){
	return &entity.Product{}, nil
}

func (q *ProductRepo) DeleteProduct(id int64) (error){
	return nil
}