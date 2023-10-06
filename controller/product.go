package controller

import (
	"format-clean-architecture/service/interface/usecase"
	"net/http"
)

type ProductController struct {
	service usecase.IProductService
}

func NewProductController(service usecase.IProductService) *ProductController{
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) CreateProduct(res http.ResponseWriter, req *http.Request){

}

func (c *ProductController) GetProduct(res http.ResponseWriter, req *http.Request){

}

func (c *ProductController) DeleteProduct(res http.ResponseWriter, req *http.Request){
	
}