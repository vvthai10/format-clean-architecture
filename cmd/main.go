package main

import (
	"fmt"
	"format-clean-architecture/api/router"
	"format-clean-architecture/config"
	"format-clean-architecture/controller"
	"format-clean-architecture/repository/postgres"
	"format-clean-architecture/service/product"
	"format-clean-architecture/service/user"
	"log"
	"net/http"
)

const (
	port string = ":8000"
)

func main() {
	config, err := config.LoadConfig("./config", "dev")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	fmt.Println(config)

	// n := negroni.New(middleware.Cors)
	// TODO: connect DB Postgres
	dbPostgres :=postgres.ConnectDB()
	// TODO: repo, service, controller, router
	userRepo := postgres.NewUserRepo(dbPostgres)
	productRepo := postgres.NewProductRepo(dbPostgres)

	userService := user.NewService(userRepo)
	productService := product.NewService(productRepo)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)

	httpRouter := router.NewGinRouter()

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and running...")
	})

	httpRouter.POST("/users", userController.CreateUser)
	httpRouter.GET("/users", userController.GetUser)

	httpRouter.POST("/products", productController.CreateProduct)
	httpRouter.GET("/products", productController.GetProduct)

	httpRouter.SERVE(port)
}