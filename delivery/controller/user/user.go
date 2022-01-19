package user

import (
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/delivery/midware"
	"rest-api/design-pattern/entity"
	userRepo "rest-api/design-pattern/repository/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{
		repository: user,
	}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		users, err := uc.repository.GetAll()

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get all users failed", nil))
		}

		if len(users) == 0 {
			return c.JSON(code, common.SimpleResponse(code, "users directory empty", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get all users success", users))
	}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		user, err := uc.repository.Get(id)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get user failed", nil))
		}

		if user == (common.UserResponse{}) {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "user does not exist", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get user success", []common.UserResponse{user}))
	}
}

func (uc UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := entity.User{}
		code := http.StatusOK

		if err := c.Bind(&user); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		id, err := uc.repository.Create(user)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "create user failed", nil))
		}

		user.Id = id

		return c.JSON(code, common.SimpleResponse(code, "create user success", []entity.User{user}))
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		user := entity.User{}

		if err := c.Bind(&user); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		user.Id = id

		if code, err := uc.repository.Update(user); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "update user success", []entity.User{user}))
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := http.StatusOK

		if valid := midware.ValidateToken(c); !valid {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid user id", nil))
		}

		if code, err := uc.repository.Delete(id); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "delete user success", nil))
	}
}
