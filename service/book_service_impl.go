package service

import (
	"context"
	"gobasic/data/request"
	"gobasic/data/response"
	"gobasic/helper"
	"gobasic/model"
	"gobasic/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookReponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookReponse

	for _, value := range books {
		book := response.BookReponse{Id: value.Id, Name: value.Name, Author: value.Author, Description: value.Description, Image: value.Image, Genre: value.Genre, PublicDate: value.PublicDate}
		bookResp = append(bookResp, book)
	}

	return bookResp
}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) (response.BookReponse, interface{}) {
	book, err := b.BookRepository.FindById(ctx, bookId)

	return response.BookReponse(book), err
}

// Create implements BookService.
func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) response.BookReponse {
	book := model.Book{
		Name:        request.Name,
		Description: request.Description,
		Author:      request.Author,
		Image:       request.Image,
		Genre:       request.Genre,
		PublicDate:  request.PublicDate,
	}
	result, err := b.BookRepository.Save(ctx, book)
	helper.PanicIfError(err)
	book.Id = result.Id

	return response.BookReponse(book)
}

// Update implements BookService.
func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) response.BookReponse {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	} else {
		book.Name = request.Name
		book.Description = request.Description
		book.Author = request.Author
		book.Genre = request.Genre
		book.PublicDate = request.PublicDate
		book.Image = request.Image
		b.BookRepository.Update(ctx, book)

	}

	return response.BookReponse(book)
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) response.BookReponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	if err != nil {
		helper.PanicIfError(err)
	} else {
		b.BookRepository.Delete(ctx, book.Id)
	}

	return response.BookReponse(book)
}
