package controller

import (
	"fmt"
	"gobasic/data/request"
	"gobasic/data/response"
	"gobasic/helper"
	"gobasic/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)

	controller.BookService.Create(requests.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}

	helper.WriteReponse(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id
	controller.BookService.Update(requests.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteReponse(writer, webResponse)
}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	controller.BookService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	helper.WriteReponse(writer, webResponse)
}

func (controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helper.WriteReponse(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	result := controller.BookService.FindById(requests.Context(), id)
	fmt.Println(result)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   result,
	}
	helper.WriteReponse(writer, webResponse)
}
