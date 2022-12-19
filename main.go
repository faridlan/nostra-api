package main

import (
	"net/http"

	"github.com/faridlan/nostra-api/app"
	"github.com/faridlan/nostra-api/controller"
	"github.com/faridlan/nostra-api/helper"
	"github.com/faridlan/nostra-api/repository"
	"github.com/faridlan/nostra-api/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	db := app.NewConnection()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	UserController := controller.NewUserController(userService)

	router.POST("/api/user", UserController.Create)
	router.PUT("/api/user/:userId", UserController.Update)
	router.DELETE("/api/user/:userId", UserController.Delete)
	router.GET("/api/user/:userId", UserController.FindById)
	router.GET("/api/user", UserController.FindAll)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
