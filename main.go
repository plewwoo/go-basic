package main

import (
	"fmt"
	"gobasic/config"
	"gobasic/controller"
	"gobasic/helper"
	"gobasic/repository"
	"gobasic/router"
	"gobasic/service"
	"net/http"
)

func main() {
	fmt.Println("Start Server")
	// Database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	//router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8000", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
