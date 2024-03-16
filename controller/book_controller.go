package controller

import (
	"gobasic/data/request"
	"gobasic/data/response"
	"gobasic/helper"
	"gobasic/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) FindAll(ctx *fiber.Ctx) error {
	result := controller.BookService.FindAll(ctx.Context())
	webResponse := response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "ok",
		Data:   result,
	}
	res := helper.WriteReponse(ctx, webResponse, fiber.StatusOK)
	return res
}

func (controller *BookController) FindById(ctx *fiber.Ctx) error {
	var res error
	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	result, errResult := controller.BookService.FindById(ctx.Context(), id)
	if (response.BookReponse{}) != result {
		webResponse := response.WebResponse{
			Code:   fiber.StatusOK,
			Status: "ok",
			Data:   result,
		}
		res = helper.WriteReponse(ctx, webResponse, fiber.StatusOK)
	} else {
		if errorMessage, ok := errResult.(map[string]string); ok {
			message := errorMessage["message"]
			webResponse := response.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: message,
				Data:   nil,
			}
			res = helper.WriteReponse(ctx, webResponse, fiber.StatusNotFound)
		}
	}
	return res
}

func (controller *BookController) Create(ctx *fiber.Ctx) error {
	bookCreateRequest := new(request.BookCreateRequest)
	helper.ReadRequestBody(ctx, &bookCreateRequest)
	result := controller.BookService.Create(ctx.Context(), *bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "ok",
		Data:   result,
	}
	res := helper.WriteReponse(ctx, webResponse, fiber.StatusCreated)
	return res
}

func (controller *BookController) Update(ctx *fiber.Ctx) error {
	var res error
	bookUpdateRequest := new(request.BookUpdateRequest)
	helper.ReadRequestBody(ctx, &bookUpdateRequest)
	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id
	result := controller.BookService.Update(ctx.Context(), *bookUpdateRequest)
	if (response.BookReponse{}) != result {
		webResponse := response.WebResponse{
			Code:   fiber.StatusOK,
			Status: "ok",
			Data:   result,
		}
		res = helper.WriteReponse(ctx, webResponse, fiber.StatusOK)
	} else {
		webResponse := response.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "not found",
			Data:   nil,
		}
		res = helper.WriteReponse(ctx, webResponse, fiber.StatusNotFound)
	}
	return res
}

func (controller *BookController) Delete(ctx *fiber.Ctx) error {
	var res error
	bookId := ctx.Params("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	result := controller.BookService.Delete(ctx.Context(), id)
	if (response.BookReponse{}) != result {
		webResponse := response.WebResponse{
			Code:   fiber.StatusOK,
			Status: "ok",
			Data:   result,
		}
		res = helper.WriteReponse(ctx, webResponse, fiber.StatusOK)
	} else {
		webResponse := response.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "not found",
			Data:   nil,
		}
		res = helper.WriteReponse(ctx, webResponse, fiber.StatusNotFound)
	}
	return res
}
