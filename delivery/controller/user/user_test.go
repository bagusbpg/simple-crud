package user

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

type mockUserRepositorySuccess struct{}

func (m mockUserRepositorySuccess) GetAll() ([]common.UserResponse, error) {
	return []common.UserResponse{
		{
			Id:    1,
			Name:  "user1",
			Email: "email1",
		},
		{
			Id:    2,
			Name:  "user2",
			Email: "email2",
		},
	}, nil
}

func (m mockUserRepositorySuccess) Get(int) (common.UserResponse, error) {
	return common.UserResponse{
		Id:    1,
		Name:  "user",
		Email: "email",
	}, nil
}

func (m mockUserRepositorySuccess) Create(entity.User) (int, error) {
	return 1, nil
}

func (m mockUserRepositorySuccess) Update(entity.User) (int, error) {
	return http.StatusOK, nil
}

func (m mockUserRepositorySuccess) Delete(int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllUsersSuccess(t *testing.T) {
	t.Run("TestGetAllUsersSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositorySuccess{})
		midware.JWTMiddleware()(userController.GetAll())(context)

		actual := common.GetAllUsersResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllUsersResponse{
			Code:    http.StatusOK,
			Message: "get all users success",
			Data: []common.UserResponse{
				{
					Id:    1,
					Name:  "user1",
					Email: "email1",
				},
				{
					Id:    2,
					Name:  "user2",
					Email: "email2",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetUserSuccess(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositorySuccess{})
		midware.JWTMiddleware()(userController.Get())(context)

		actual := common.GetUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetUserResponse{
			Code:    http.StatusOK,
			Message: "get user success",
			Data: []common.UserResponse{
				{
					Id:    1,
					Name:  "user",
					Email: "email",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateUserSuccess(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user",
			"email":    "email",
			"password": "password",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositorySuccess{})
		userController.Create()(context)

		actual := common.CreateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateUserResponse{
			Code:    http.StatusOK,
			Message: "create user success",
			Data: []entity.User{
				{
					Id:       1,
					Name:     "user",
					Email:    "email",
					Password: "password",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateUserSuccess(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user",
			"email":    "email",
			"password": "password",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositorySuccess{})
		midware.JWTMiddleware()(userController.Update())(context)

		actual := common.UpdateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateUserResponse{
			Code:    http.StatusOK,
			Message: "update user success",
			Data: []entity.User{
				{
					Id:       1,
					Name:     "user",
					Email:    "email",
					Password: "password",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteUserSuccess(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositorySuccess{})
		midware.JWTMiddleware()(userController.Delete())(context)

		actual := common.DeleteUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteUserResponse{
			Code:    http.StatusOK,
			Message: "delete user success",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN REPOSITORY

type mockUserRepositoryFailRepo struct{}

func (m mockUserRepositoryFailRepo) GetAll() ([]common.UserResponse, error) {
	return nil, assert.AnError
}

func (m mockUserRepositoryFailRepo) Get(int) (common.UserResponse, error) {
	return common.UserResponse{}, assert.AnError
}

func (m mockUserRepositoryFailRepo) Create(entity.User) (int, error) {
	return 0, fmt.Errorf("create user failed")
}

func (m mockUserRepositoryFailRepo) Update(entity.User) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("update user failed")
}

func (m mockUserRepositoryFailRepo) Delete(int) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("delete user failed")
}

func TestGetAllUsersFailRepo(t *testing.T) {
	t.Run("TestGetAllUsersFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositoryFailRepo{})
		midware.JWTMiddleware()(userController.GetAll())(context)

		actual := common.GetAllUsersResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllUsersResponse{
			Code:    http.StatusInternalServerError,
			Message: "get all users failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetUserFailRepo(t *testing.T) {
	t.Run("TestGetUserFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositoryFailRepo{})
		midware.JWTMiddleware()(userController.Get())(context)

		actual := common.GetUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetUserResponse{
			Code:    http.StatusInternalServerError,
			Message: "get user failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateUserFailRepo(t *testing.T) {
	t.Run("TestCreateUserFailRepo", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user",
			"email":    "email",
			"password": "password",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositoryFailRepo{})
		userController.Create()(context)

		actual := common.CreateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateUserResponse{
			Code:    http.StatusInternalServerError,
			Message: "create user failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateUserFailRepo(t *testing.T) {
	t.Run("TestUpdateUserFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user",
			"email":    "email",
			"password": "password",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositoryFailRepo{})
		midware.JWTMiddleware()(userController.Update())(context)

		actual := common.UpdateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateUserResponse{
			Code:    http.StatusInternalServerError,
			Message: "update user failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteUserFailRepo(t *testing.T) {
	t.Run("TestDeleteUserFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositoryFailRepo{})
		midware.JWTMiddleware()(userController.Delete())(context)

		actual := common.DeleteUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteUserResponse{
			Code:    http.StatusInternalServerError,
			Message: "delete user failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN BINDING, INVALID ID, OR EMPTY DIRECTORY

type mockUserRepositoryFailOther struct{}

func (m mockUserRepositoryFailOther) GetAll() ([]common.UserResponse, error) {
	return []common.UserResponse{}, nil
}

func (m mockUserRepositoryFailOther) Get(int) (common.UserResponse, error) {
	return common.UserResponse{}, nil
}

func (m mockUserRepositoryFailOther) Create(entity.User) (int, error) {
	return 0, nil
}

func (m mockUserRepositoryFailOther) Update(entity.User) (int, error) {
	return http.StatusOK, nil
}

func (m mockUserRepositoryFailOther) Delete(int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllUsersEmptyDirectory(t *testing.T) {
	t.Run("TestGetAllUsersEmptyDirectory", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositoryFailOther{})
		midware.JWTMiddleware()(userController.GetAll())(context)

		actual := common.GetAllUsersResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllUsersResponse{
			Code:    http.StatusOK,
			Message: "users directory empty",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetUserFailInvalidId(t *testing.T) {
	t.Run("TestGetUserFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		userController := New(mockUserRepositoryFailOther{})
		midware.JWTMiddleware()(userController.Get())(context)

		actual := common.GetUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetUserResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetUserDoesNotExist(t *testing.T) {
	t.Run("TestGetUserDoesNotExist", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositoryFailOther{})
		midware.JWTMiddleware()(userController.Get())(context)

		actual := common.GetUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetUserResponse{
			Code:    http.StatusBadRequest,
			Message: "user does not exist",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateUserFailBinding(t *testing.T) {
	t.Run("TestCreateUserFailBinding", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "user",
			"email":    "email",
			"password": 1234,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockUserRepositoryFailOther{})
		userController.Create()(context)

		actual := common.CreateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateUserResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateUserFailInvalidId(t *testing.T) {
	t.Run("TestUpdateUserFailBinding", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "user",
			"email":    "email",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		userController := New(mockUserRepositoryFailOther{})
		midware.JWTMiddleware()(userController.Update())(context)

		actual := common.UpdateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateUserResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateUserFailBinding(t *testing.T) {
	t.Run("TestUpdateUserFailBinding", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "user",
			"email":    "email",
			"password": 1234,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepositoryFailOther{})
		midware.JWTMiddleware()(userController.Update())(context)

		actual := common.UpdateUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateUserResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteUserFailInvalidId(t *testing.T) {
	t.Run("TestDeleteUserFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		userController := New(mockUserRepositoryFailRepo{})
		midware.JWTMiddleware()(userController.Delete())(context)

		actual := common.DeleteUserResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteUserResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}
