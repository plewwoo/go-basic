package service

import (
	"context"
	"gobasic/data/request"
	"gobasic/data/response"
)

type BookService interface {
	FindAll(ctx context.Context) []response.BookReponse
	FindById(ctx context.Context, bookId int) (response.BookReponse, interface{})
	Create(ctx context.Context, request request.BookCreateRequest) response.BookReponse
	Update(ctx context.Context, request request.BookUpdateRequest) response.BookReponse
	Delete(ctx context.Context, bookId int) response.BookReponse
}
