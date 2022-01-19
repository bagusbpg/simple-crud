package book

import (
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/delivery/midware"
	"rest-api/design-pattern/entity"
	bookRepo "rest-api/design-pattern/repository/book"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	repository bookRepo.Book
}

func New(book bookRepo.Book) *BookController {
	return &BookController{
		repository: book,
	}
}

func (bc BookController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK
		books, err := bc.repository.GetAll()

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get all books failed", nil))
		}

		if len(books) == 0 {
			return c.JSON(code, common.SimpleResponse(code, "books directory empty", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get all books success", books))
	}
}

func (bc BookController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid book id", nil))
		}

		book, err := bc.repository.Get(id)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get book failed", nil))
		}

		if book == (common.BookResponse{}) {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "book does not exist", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get book success", []common.BookResponse{book}))
	}
}

func (bc BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		book := entity.Book{}
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		if err := c.Bind(&book); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		id, err := bc.repository.Create(book)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "create book failed", nil))
		}

		book.Id = id

		return c.JSON(code, common.SimpleResponse(code, "create book success", []entity.Book{book}))
	}
}

func (bc BookController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid book id", nil))
		}

		book := entity.Book{}

		if err := c.Bind(&book); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		book.Id = id

		if code, err := bc.repository.Update(book); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "update book success", []entity.Book{book}))
	}
}

func (bc BookController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid book id", nil))
		}

		if code, err := bc.repository.Delete(id); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "delete book success", nil))
	}
}
