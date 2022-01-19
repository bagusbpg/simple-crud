package auth

import (
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
	authRepo "rest-api/design-pattern/repository/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repository authRepo.Auth
}

func New(auth authRepo.Auth) *AuthController {
	return &AuthController{
		repository: auth,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		login := entity.User{}

		if err := c.Bind(&login); err != nil {
			code := http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", ""))
		}

		token, code := a.repository.Login(login.Name, login.Password)

		if code != http.StatusOK {
			return c.JSON(code, common.SimpleResponse(code, token, ""))
		}

		return c.JSON(code, common.SimpleResponse(code, "login success", token))
	}
}
