package main

import (
	"rest-api/design-pattern/config"

	_authController "rest-api/design-pattern/delivery/controller/auth"
	_bookController "rest-api/design-pattern/delivery/controller/book"
	_productController "rest-api/design-pattern/delivery/controller/product"
	_userController "rest-api/design-pattern/delivery/controller/user"
	"rest-api/design-pattern/delivery/midware"
	"rest-api/design-pattern/delivery/router"

	_authRepo "rest-api/design-pattern/repository/auth"
	_bookRepo "rest-api/design-pattern/repository/book"
	_productRepo "rest-api/design-pattern/repository/product"
	_userRepo "rest-api/design-pattern/repository/user"
	"rest-api/design-pattern/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fallback := config.AppConfig{Type: "main"}
	config := config.GetConfig(&fallback)

	db := util.GetDBInstance(config)
	defer db.Close()

	authRepo := _authRepo.New(db)
	bookRepo := _bookRepo.New(db)
	productRepo := _productRepo.New(db)
	userRepo := _userRepo.New(db)

	authController := _authController.New(authRepo)
	bookController := _bookController.New(bookRepo)
	productController := _productController.New(productRepo)
	userController := _userController.New(userRepo)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash(), midware.CustomLogger())

	router.RegisterPath(e, authController, bookController, userController, productController)

	e.Logger.Fatal(e.Start((":8080")))
}
