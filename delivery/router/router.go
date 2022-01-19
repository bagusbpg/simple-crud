package router

import (
	"rest-api/design-pattern/delivery/controller/auth"
	"rest-api/design-pattern/delivery/controller/book"
	"rest-api/design-pattern/delivery/controller/product"
	"rest-api/design-pattern/delivery/controller/user"
	"rest-api/design-pattern/delivery/midware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	authController *auth.AuthController,
	bookController *book.BookController,
	userController *user.UserController,
	productController *product.ProductController,
) {

	// Login
	e.POST("/login", authController.Login())

	// User
	e.GET("/users", userController.GetAll(), midware.JWTMiddleware())
	e.GET("/users/:id", userController.Get(), midware.JWTMiddleware())
	e.POST("/users", userController.Create())
	e.PUT("/users/:id", userController.Update(), midware.JWTMiddleware())
	e.DELETE("/users/:id", userController.Delete(), midware.JWTMiddleware())

	// Book
	e.GET("/books", bookController.GetAll())
	e.GET("/books/:id", bookController.Get())
	e.POST("/books", bookController.Create(), midware.JWTMiddleware())
	e.PUT("/books/:id", bookController.Update(), midware.JWTMiddleware())
	e.DELETE("/books/:id", bookController.Delete(), midware.JWTMiddleware())

	// Product
	e.GET("/products", productController.GetAll())
	e.GET("/products/:id", productController.Get())
	e.POST("/products", productController.Create(), midware.JWTMiddleware())
	e.PUT("/products/:id", productController.Update(), midware.JWTMiddleware())
	e.DELETE("/products/:id", productController.Delete(), midware.JWTMiddleware())
}
