package resolver

import (
	"context"
	"errors"
	"time"

	"github.com/Ikhlashmulya/golang-graphql-book/model"
	"github.com/Ikhlashmulya/golang-graphql-book/service"

	"github.com/graphql-go/graphql"
)

var (
	bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})

	RootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)

					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					response, err := bookService.FindById(timeoutCtx, id)
					return response, err
				},
			},
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					responses, err := bookService.FindAll(timeoutCtx)
					return responses, err
				},
			},
			"getBookByTitle": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title := params.Args["title"].(string)

					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					response, err := bookService.FindByTitle(timeoutCtx, title)
					return response, err
				},
			},
		},
	})

	bookInput = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "BookInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})

	RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"createBook": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(bookInput),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					input := params.Args["input"].(map[string]any)

					title := input["title"].(string)
					description := input["description"].(string)
					author := input["author"].(string)

					creteBookRequest := model.CreateBookRequest{
						Title:       title,
						Description: description,
						Author:      author,
					}

					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					response, err := bookService.Create(timeoutCtx, creteBookRequest)
					return response, err
				},
			},
			"updateBook": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(bookInput),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if !ok {
						return nil, errors.New("invalid argument type for 'id'")
					}

					input := params.Args["input"].(map[string]any)

					title := input["title"].(string)
					description := input["description"].(string)
					author := input["author"].(string)

					// input, ok := params.Args["input"].(map[string]any)
					// if !ok {
					// 	return nil, errors.New("invalid argument type for 'input'")
					// }

					// title, ok := input["title"].(string)
					// if !ok {
					// 	return nil, errors.New("invalid argument type for 'input'")
					// }

					// description, ok := input["description"].(string)
					// if !ok {
					// 	return nil, errors.New("invalid argument type for 'input'")
					// }

					// author, ok := input["author"].(string)
					// if !ok {
					// 	return nil, errors.New("invalid argument type for 'input'")
					// }

					updateBookRequest := model.UpdateBookRequest{
						Id:          id,
						Title:       title,
						Description: description,
						Author:      author,
					}

					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					response, err := bookService.Update(timeoutCtx, updateBookRequest)
					return response, err
				},
			},
			"deleteBook": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if !ok {
						return nil, errors.New("invalid argument type for 'id'")
					}

					timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					bookService := params.Context.Value("bookService").(service.BookService)
					response, err := bookService.Delete(timeoutCtx, id)
					return response, err
				},
			},
		},
	})
)
