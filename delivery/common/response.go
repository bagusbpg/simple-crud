package common

import "rest-api/design-pattern/entity"

type ProductResponse struct {
	Id       int    `json:"id" form:"id"`
	Merchant string `json:"merchant" form:"merchant"`
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
}

type BookResponse struct {
	Id        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Language  string `json:"language" form:"language"`
	Pages     int    `json:"page" form:"page"`
	ISBN13    string `json:"isbn13" form:"isbn13"`
}

type UserResponse struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type GetAllUsersResponse struct {
	Code    int            `json:"code" form:"code"`
	Message string         `json:"message" form:"message"`
	Data    []UserResponse `json:"data" form:"data"`
}

type GetUserResponse struct {
	Code    int            `json:"code" form:"code"`
	Message string         `json:"message" form:"message"`
	Data    []UserResponse `json:"data" form:"data"`
}

type CreateUserResponse struct {
	Code    int           `json:"code" form:"code"`
	Message string        `json:"message" form:"message"`
	Data    []entity.User `json:"data" form:"data"`
}

type UpdateUserResponse struct {
	Code    int           `json:"code" form:"code"`
	Message string        `json:"message" form:"message"`
	Data    []entity.User `json:"data" form:"data"`
}

type DeleteUserResponse struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

type GetAllProductsResponse struct {
	Code    int               `json:"code" form:"code"`
	Message string            `json:"message" form:"message"`
	Data    []ProductResponse `json:"data" form:"data"`
}

type GetProductResponse struct {
	Code    int               `json:"code" form:"code"`
	Message string            `json:"message" form:"message"`
	Data    []ProductResponse `json:"data" form:"data"`
}

type CreateProductResponse struct {
	Code    int               `json:"code" form:"code"`
	Message string            `json:"message" form:"message"`
	Data    []ProductResponse `json:"data" form:"data"`
}

type UpdateProductResponse struct {
	Code    int              `json:"code" form:"code"`
	Message string           `json:"message" form:"message"`
	Data    []entity.Product `json:"data" form:"data"`
}

type DeleteProductResponse struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

type GetAllBooksResponse struct {
	Code    int            `json:"code" form:"code"`
	Message string         `json:"message" form:"message"`
	Data    []BookResponse `json:"data" form:"data"`
}

type GetBookResponse struct {
	Code    int            `json:"code" form:"code"`
	Message string         `json:"message" form:"message"`
	Data    []BookResponse `json:"data" form:"data"`
}

type CreateBookResponse struct {
	Code    int           `json:"code" form:"code"`
	Message string        `json:"message" form:"message"`
	Data    []entity.Book `json:"data" form:"data"`
}

type UpdateBookResponse struct {
	Code    int           `json:"code" form:"code"`
	Message string        `json:"message" form:"message"`
	Data    []entity.Book `json:"data" form:"data"`
}

type DeleteBookResponse struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

type LoginResponse struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}
