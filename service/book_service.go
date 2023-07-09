package service

import (
	"github.com/Ikhlashmulya/golang-graphql-book/entity"
	"github.com/Ikhlashmulya/golang-graphql-book/mapper"
	"github.com/Ikhlashmulya/golang-graphql-book/model"
	"github.com/Ikhlashmulya/golang-graphql-book/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type BookService interface {
	Create(ctx context.Context, request model.CreateBookRequest) (response model.BookResponse, err error)
	Update(ctx context.Context, request model.UpdateBookRequest) (response model.BookResponse, err error)
	Delete(ctx context.Context, bookId string) (response model.BookResponse, err error)
	FindById(ctx context.Context, bookId string) (response model.BookResponse, err error)
	FindByTitle(ctx context.Context, bookTitle string) (response model.BookResponse, err error)
	FindAll(ctx context.Context) (responses []model.BookResponse, err error)
}

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

func (service *BookServiceImpl) Create(ctx context.Context, request model.CreateBookRequest) (response model.BookResponse, err error) {
	book := entity.Book{
		Id:          uuid.NewString(),
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
	}

	err = service.BookRepository.Create(ctx, book)
	response = mapper.ToBookResponse(&book)

	return response, err
}

func (service *BookServiceImpl) Update(ctx context.Context, request model.UpdateBookRequest) (response model.BookResponse, err error) {
	book, err := service.BookRepository.FindById(ctx, request.Id)
	if err != nil {
		return model.BookResponse{}, err
	}

	if request.Title != "" {
		book.Title = request.Title
	}

	if request.Author != "" {
		book.Author = request.Author
	}

	if request.Description != "" {
		book.Description = request.Description
	}

	err = service.BookRepository.Update(ctx, book)
	response = mapper.ToBookResponse(&book)

	return response, err
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId string) (response model.BookResponse, err error) {
	book, err := service.BookRepository.FindById(ctx, bookId)
	if err != nil {
		return model.BookResponse{}, err
	}

	err = service.BookRepository.Delete(ctx, bookId)
	response = mapper.ToBookResponse(&book)
	return response, err
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId string) (response model.BookResponse, err error) {
	book, err := service.BookRepository.FindById(ctx, bookId)
	response = mapper.ToBookResponse(&book)
	return response, err
}

func (service *BookServiceImpl) FindByTitle(ctx context.Context, bookTitle string) (response model.BookResponse, err error) {
	book, err := service.BookRepository.FindByTitle(ctx, bookTitle)
	response = mapper.ToBookResponse(&book)
	return response, err
}

func (service *BookServiceImpl) FindAll(ctx context.Context) (responses []model.BookResponse, err error) {
	books := service.BookRepository.FindAll(ctx)
	if len(books) == 0 {
		err = errors.New("no content")
	}

	for _, book := range books {
		responses = append(responses, mapper.ToBookResponse(&book))
	}

	return responses, err
}
