package handler

import (
	"context"

	"github.com/Ikhlashmulya/golang-graphql-book/graph/resolver"
	"github.com/Ikhlashmulya/golang-graphql-book/service"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{BookService: bookService}
}

func (bookHandler *BookHandler) Graph(ctx *fiber.Ctx) error {
	var query map[string]any
	ctx.BodyParser(&query)

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    resolver.RootQuery,
		Mutation: resolver.RootMutation,
	})

	params := graphql.Params{
		Schema:        schema,
		RequestString: query["query"].(string),
		Context:       context.WithValue(context.Background(), "bookService", bookHandler.BookService),
	}

	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		ctx.SendString("failed to execute graphql operation")
	}

	return ctx.JSON(result)
}
