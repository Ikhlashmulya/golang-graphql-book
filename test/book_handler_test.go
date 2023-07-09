package test

import (
	"github.com/Ikhlashmulya/golang-graphql-book/config"
	"github.com/Ikhlashmulya/golang-graphql-book/graph/handler"
	"github.com/Ikhlashmulya/golang-graphql-book/repository"
	"github.com/Ikhlashmulya/golang-graphql-book/service"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var configuration = config.NewConfig("../.env.test")
var db = config.NewDB(configuration)
var bookRepository = repository.NewBookRepository(db)
var bookService = service.NewBookService(bookRepository)
var bookHandler = handler.NewBookHandler(bookService)
var app = setupTestApp()

func setupTestApp() *fiber.App {
	app := fiber.New()

	app.Post("/", bookHandler.Graph)

	return app
}
func TestMutation(t *testing.T) {

	requestBody := `{"query": "mutation { createBook(input: {title:\"ini title\", author:\"ini author\", description:\"ini description\"}) { id } }"}`
	request := httptest.NewRequest(fiber.MethodPost, "/", strings.NewReader(requestBody))
	request.Header.Add("content-type", "application/json")

	result, err := app.Test(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, result.StatusCode)

	bytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	assert.NotNil(t, responseBody["data"].(map[string]any)["createBook"])
}

func TestQuery(t *testing.T) {

	requestBody := `{"query": "query { books { id, title } }"}`
	request := httptest.NewRequest(fiber.MethodPost, "/", strings.NewReader(requestBody))
	request.Header.Add("content-type", "application/json")

	result, err := app.Test(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, result.StatusCode)

	bytes, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	assert.NotNil(t, responseBody["data"])
}

func TestMain(m *testing.M) {
	d, _ := db.DB()
	d.Exec("truncate books")
	m.Run()
	d.Exec("truncate books")
}
