package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/delivery/midware"
	"rest-api/design-pattern/entity"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TEST SUCCESS

type mockBookRepositorySuccess struct{}

func (m mockBookRepositorySuccess) GetAll() ([]common.BookResponse, error) {
	return []common.BookResponse{
		{
			Id:        1,
			Title:     "title1",
			Author:    "author1",
			Publisher: "publisher1",
			Language:  "language1",
			Pages:     100,
			ISBN13:    "isbn1",
		},
		{
			Id:        2,
			Title:     "title2",
			Author:    "author2",
			Publisher: "publisher2",
			Language:  "language2",
			Pages:     100,
			ISBN13:    "isbn2",
		},
	}, nil
}

func (m mockBookRepositorySuccess) Get(int) (common.BookResponse, error) {
	return common.BookResponse{
		Id:        1,
		Title:     "title1",
		Author:    "author1",
		Publisher: "publisher1",
		Language:  "language1",
		Pages:     100,
		ISBN13:    "isbn1",
	}, nil
}

func (m mockBookRepositorySuccess) Create(entity.Book) (int, error) {
	return 1, nil
}

func (m mockBookRepositorySuccess) Update(entity.Book) (int, error) {
	return http.StatusOK, nil
}

func (m mockBookRepositorySuccess) Delete(int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllBooksSuccess(t *testing.T) {
	t.Run("TestGetAllBooksSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositorySuccess{})
		bookController.GetAll()(context)

		actual := common.GetAllBooksResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllBooksResponse{
			Code:    http.StatusOK,
			Message: "get all books success",
			Data: []common.BookResponse{
				{
					Id:        1,
					Title:     "title1",
					Author:    "author1",
					Publisher: "publisher1",
					Language:  "language1",
					Pages:     100,
					ISBN13:    "isbn1",
				},
				{
					Id:        2,
					Title:     "title2",
					Author:    "author2",
					Publisher: "publisher2",
					Language:  "language2",
					Pages:     100,
					ISBN13:    "isbn2",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetBookSuccess(t *testing.T) {
	t.Run("TestGetBookSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositorySuccess{})
		bookController.Get()(context)

		actual := common.GetBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetBookResponse{
			Code:    http.StatusOK,
			Message: "get book success",
			Data: []common.BookResponse{
				{
					Id:        1,
					Title:     "title1",
					Author:    "author1",
					Publisher: "publisher1",
					Language:  "language1",
					Pages:     100,
					ISBN13:    "isbn1",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateBookSuccess(t *testing.T) {
	t.Run("TestCreateBookSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     100,
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositorySuccess{})
		midware.JWTMiddleware()(bookController.Create())(context)

		actual := common.CreateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateBookResponse{
			Code:    http.StatusOK,
			Message: "create book success",
			Data: []entity.Book{
				{
					Id:        1,
					Title:     "title1",
					Author:    "author1",
					Publisher: "publisher1",
					Language:  "language1",
					Pages:     100,
					ISBN13:    "isbn1",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateBookSuccess(t *testing.T) {
	t.Run("TestUpdateBookSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     100,
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositorySuccess{})
		midware.JWTMiddleware()(bookController.Update())(context)

		actual := common.UpdateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateBookResponse{
			Code:    http.StatusOK,
			Message: "update book success",
			Data: []entity.Book{
				{
					Id:        1,
					Title:     "title1",
					Author:    "author1",
					Publisher: "publisher1",
					Language:  "language1",
					Pages:     100,
					ISBN13:    "isbn1",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteBookSuccess(t *testing.T) {
	t.Run("TestDeleteBookSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositorySuccess{})
		midware.JWTMiddleware()(bookController.Delete())(context)

		actual := common.DeleteBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteBookResponse{
			Code:    http.StatusOK,
			Message: "delete book success",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN REPOSITORY

type mockBookRepositoryFailRepo struct{}

func (m mockBookRepositoryFailRepo) GetAll() ([]common.BookResponse, error) {
	return nil, assert.AnError
}

func (m mockBookRepositoryFailRepo) Get(int) (common.BookResponse, error) {
	return common.BookResponse{}, assert.AnError
}

func (m mockBookRepositoryFailRepo) Create(entity.Book) (int, error) {
	return 0, assert.AnError
}

func (m mockBookRepositoryFailRepo) Update(entity.Book) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("udate book failed")
}

func (m mockBookRepositoryFailRepo) Delete(int) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("delete book failed")
}

func TestGetAllBooksFailRepo(t *testing.T) {
	t.Run("TestGetAllBooksFailRepo", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositoryFailRepo{})
		bookController.GetAll()(context)

		actual := common.GetAllBooksResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllBooksResponse{
			Code:    http.StatusInternalServerError,
			Message: "get all books failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetBookFailRepo(t *testing.T) {
	t.Run("TestGetBookFailRepo", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositoryFailRepo{})
		bookController.Get()(context)

		actual := common.GetBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetBookResponse{
			Code:    http.StatusInternalServerError,
			Message: "get book failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateBookFailRepo(t *testing.T) {
	t.Run("TestCreateBookFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     100,
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositoryFailRepo{})
		midware.JWTMiddleware()(bookController.Create())(context)

		actual := common.CreateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateBookResponse{
			Code:    http.StatusInternalServerError,
			Message: "create book failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateBookFailRepo(t *testing.T) {
	t.Run("TestUpdateBookFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     100,
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositoryFailRepo{})
		midware.JWTMiddleware()(bookController.Update())(context)

		actual := common.UpdateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateBookResponse{
			Code:    http.StatusInternalServerError,
			Message: "udate book failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteBookFailRepo(t *testing.T) {
	t.Run("TestDeleteBookFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositoryFailRepo{})
		midware.JWTMiddleware()(bookController.Delete())(context)

		actual := common.DeleteBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteBookResponse{
			Code:    http.StatusInternalServerError,
			Message: "delete book failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN BINDING, INVALID ID, OR EMPTY DIRECTORY

type mockBookRepositoryFailOther struct{}

func (m mockBookRepositoryFailOther) GetAll() ([]common.BookResponse, error) {
	return []common.BookResponse{}, nil
}

func (m mockBookRepositoryFailOther) Get(int) (common.BookResponse, error) {
	return common.BookResponse{}, nil
}

func (m mockBookRepositoryFailOther) Create(entity.Book) (int, error) {
	return 0, nil
}

func (m mockBookRepositoryFailOther) Update(entity.Book) (int, error) {
	return http.StatusOK, nil
}

func (m mockBookRepositoryFailOther) Delete(int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllBooksEmptyDirectory(t *testing.T) {
	t.Run("TestGetAllBooksEmptyDirectory", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositoryFailOther{})
		bookController.GetAll()(context)

		actual := common.GetAllBooksResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllBooksResponse{
			Code:    http.StatusOK,
			Message: "books directory empty",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetBookFailInvalidId(t *testing.T) {
	t.Run("TestGetBookFailInvalidId", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		bookController := New(mockBookRepositoryFailOther{})
		bookController.Get()(context)

		actual := common.GetBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetBookResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid book id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetBookDoesNotExist(t *testing.T) {
	t.Run("TestGetBookDoesNotExist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositoryFailOther{})
		bookController.Get()(context)

		actual := common.GetBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetBookResponse{
			Code:    http.StatusBadRequest,
			Message: "book does not exist",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateBookFailBinding(t *testing.T) {
	t.Run("TestCreateBookFailBinding", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     "100",
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books")

		bookController := New(mockBookRepositoryFailOther{})
		midware.JWTMiddleware()(bookController.Create())(context)

		actual := common.CreateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateBookResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateBookFailInvalidId(t *testing.T) {
	t.Run("TestUpdateBookFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     100,
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		bookController := New(mockBookRepositoryFailOther{})
		midware.JWTMiddleware()(bookController.Update())(context)

		actual := common.UpdateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateBookResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid book id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateBookFailBinding(t *testing.T) {
	t.Run("TestUpdateBookFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"title":     "title1",
			"author":    "author1",
			"publisher": "publisher1",
			"language":  "language1",
			"pages":     "100",
			"isbn13":    "isbn1",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		bookController := New(mockBookRepositoryFailOther{})
		midware.JWTMiddleware()(bookController.Update())(context)

		actual := common.UpdateBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateBookResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteBookFailInvalidId(t *testing.T) {
	t.Run("TestDeleteBookFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/books/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		bookController := New(mockBookRepositoryFailRepo{})
		midware.JWTMiddleware()(bookController.Delete())(context)

		actual := common.DeleteBookResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteBookResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid book id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}
