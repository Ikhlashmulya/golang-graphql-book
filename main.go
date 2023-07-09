package main

import (
	"log"

	"github.com/Ikhlashmulya/golang-graphql-book/config"
	"github.com/Ikhlashmulya/golang-graphql-book/graph/handler"
	"github.com/Ikhlashmulya/golang-graphql-book/repository"
	"github.com/Ikhlashmulya/golang-graphql-book/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	configuration := config.NewConfig()
	db := config.NewDB(configuration)
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	app := fiber.New()

	app.Post("/", bookHandler.Graph)

	err := app.Listen(":8080")
	log.Fatal(err)
}
