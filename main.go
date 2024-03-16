package main

import (
	"fmt"
	"gobasic/config"
	"gobasic/controller"
	"gobasic/helper"
	"gobasic/repository"
	"gobasic/router"
	"gobasic/service"
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

	err := routes.Listen("localhost:8000")
	helper.PanicIfError(err)
}
