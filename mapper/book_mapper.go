package mapper

import (
	"github.com/Ikhlashmulya/golang-graphql-book/entity"
	"github.com/Ikhlashmulya/golang-graphql-book/model"
)

func ToBookResponse(book *entity.Book) model.BookResponse {
	return model.BookResponse{
		Id:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}
}
