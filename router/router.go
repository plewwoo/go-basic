package router

import (
	"gobasic/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRouter(bookController *controller.BookController) *fiber.App {
	router := fiber.New()
	// Middleware
	router.Use(cors.New())
	router.Use(logger.New())

	// HTTP Method
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome Fiber")
	})
	router.Get("/api/book", bookController.FindAll)
	router.Get("/api/book/:bookId", bookController.FindById)
	router.Post("/api/book", bookController.Create)
	router.Put("/api/book/:bookId", bookController.Update)
	router.Delete("/api/book/:bookId", bookController.Delete)

	return router
}
