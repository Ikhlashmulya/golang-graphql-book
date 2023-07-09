package repository

import (
	"github.com/Ikhlashmulya/golang-graphql-book/entity"
	"context"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, book entity.Book) error
	Update(ctx context.Context, book entity.Book) error
	Delete(ctx context.Context, bookId string) error
	FindById(ctx context.Context, bookId string) (response entity.Book, err error)
	FindByTitle(ctx context.Context, bookTitle string) (response entity.Book, err error)
	FindAll(ctx context.Context) (books []entity.Book)
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (repository *BookRepositoryImpl) Create(ctx context.Context, book entity.Book) error {
	return repository.DB.WithContext(ctx).Create(&book).Error
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, book entity.Book) error {
	return repository.DB.WithContext(ctx).Save(&book).Error
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, bookId string) error {
	return repository.DB.WithContext(ctx).Delete(&entity.Book{}, "id = ?", bookId).Error
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, bookId string) (response entity.Book, err error) {
	err = repository.DB.WithContext(ctx).First(&response, "id = ?", bookId).Error
	return
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context) (books []entity.Book) {
	repository.DB.WithContext(ctx).Find(&books)
	return
}

func (repository *BookRepositoryImpl) FindByTitle(ctx context.Context, bookTitle string) (response entity.Book, err error) {
	err = repository.DB.WithContext(ctx).First(&response, "title LIKE ?", "%"+bookTitle+"%").Error
	return
}