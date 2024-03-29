package repository

import (
	"context"
	"database/sql"
	"gobasic/helper"
	"gobasic/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT * FROM book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name, &book.Description, &book.Author, &book.Image, &book.Genre, &book.PublicDate)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, interface{}) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT * FROM book WHERE id = ?"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name, &book.Description, &book.Author, &book.Image, &book.Genre, &book.PublicDate)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, map[string]string{"message": "book id not found"}
	}
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) (model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO book (name, description, author, image, genre, public_date) VALUES (?, ?, ?, ?, ?, ?)"
	result, errExec := tx.ExecContext(ctx, SQL, book.Name, book.Description, book.Author, book.Image, book.Genre, book.PublicDate)
	helper.PanicIfError(errExec)
	lastInsertId, lastInsertIdErr := result.LastInsertId()
	helper.PanicIfError(lastInsertIdErr)
	book.Id = int(lastInsertId)
	return book, nil
}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE book SET name = ?, description = ?, author = ?, image = ?, genre = ?, public_date = ? WHERE id = ?"
	_, errExec := tx.ExecContext(ctx, SQL, book.Name, book.Description, book.Author, book.Image, book.Genre, book.PublicDate, book.Id)
	helper.PanicIfError(errExec)
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM book WHERE id = ?"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)
}
