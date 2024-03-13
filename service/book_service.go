package service

import (
	"context"
	"gobasic/data/request"
	"gobasic/data/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookReponse
	FindAll(ctx context.Context) []response.BookReponse
}
