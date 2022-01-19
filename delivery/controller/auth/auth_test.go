package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-api/design-pattern/delivery/common"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TEST SUCCESS

type mockAuthRepositorySuccess struct{}

func (m mockAuthRepositorySuccess) Login(string, string) (string, int) {
	return "aValidToken", http.StatusOK
}

func TestLoginSuccess(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user1",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositorySuccess{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusOK,
			Message: "login success",
			Data:    "aValidToken",
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL

type mockAuthRepositoryFailRepo struct{}

func (m mockAuthRepositoryFailRepo) Login(string, string) (string, int) {
	return "get user failed", http.StatusInternalServerError
}

func TestLoginFailRepo(t *testing.T) {
	t.Run("TestLoginFailRepo", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user1",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositoryFailRepo{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusInternalServerError,
			Message: "get user failed",
			Data:    "",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestLoginFailBinding(t *testing.T) {
	t.Run("TestLoginFailBinding", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "user1",
			"password": 1234,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositoryFailRepo{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    "",
		}

		assert.Equal(t, expected, actual)
	})
}

type mockAuthRepositoryFailUserNotFound struct{}

func (m mockAuthRepositoryFailUserNotFound) Login(string, string) (string, int) {
	return "user does not exist", http.StatusUnauthorized
}

func TestLoginFailUserNotFound(t *testing.T) {
	t.Run("TestLoginFailUserNotFound", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user1",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositoryFailUserNotFound{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusUnauthorized,
			Message: "user does not exist",
			Data:    "",
		}

		assert.Equal(t, expected, actual)
	})
}

type mockAuthRepositoryFailPasswordIncorrect struct{}

func (m mockAuthRepositoryFailPasswordIncorrect) Login(string, string) (string, int) {
	return "password incorrect", http.StatusUnauthorized
}

func TestLoginFailPasswordIncorrect(t *testing.T) {
	t.Run("TestLoginFailPasswordIncorrect", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user1",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositoryFailPasswordIncorrect{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusUnauthorized,
			Message: "password incorrect",
			Data:    "",
		}

		assert.Equal(t, expected, actual)
	})
}

type mockAuthRepositoryFailTokenCreation struct{}

func (m mockAuthRepositoryFailTokenCreation) Login(string, string) (string, int) {
	return "token creation failed", http.StatusInternalServerError
}

func TestLoginFailTokenCreation(t *testing.T) {
	t.Run("TestLoginFailTokenCreation", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "user1",
			"password": "password1",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		authController := New(mockAuthRepositoryFailTokenCreation{})
		authController.Login()(context)

		actual := common.LoginResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.LoginResponse{
			Code:    http.StatusInternalServerError,
			Message: "token creation failed",
			Data:    "",
		}

		assert.Equal(t, expected, actual)
	})
}
