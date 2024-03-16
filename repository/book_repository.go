package repository

import (
	"context"
	"gobasic/model"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book) (model.Book, error)
	Update(ctx context.Context, book model.Book)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (model.Book, interface{})
	FindAll(ctx context.Context) []model.Book
}
