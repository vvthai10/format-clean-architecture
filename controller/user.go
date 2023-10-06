package controller

import (
	"format-clean-architecture/service/interface/usecase"
	"net/http"
)

type UserController struct {
	service usecase.IUserService
}

func NewUserController(service usecase.IUserService) *UserController{
	return &UserController{
		service: service,
	}
}

func (c *UserController) CreateUser(res http.ResponseWriter, req *http.Request){

}

func (c *UserController) GetUser(res http.ResponseWriter, req *http.Request){

}

func (c *UserController) DeleteUser(res http.ResponseWriter, req *http.Request){
	
}