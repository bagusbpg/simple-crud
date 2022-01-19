package product

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

type mockProductRepositorySuccess struct{}

func (m mockProductRepositorySuccess) GetAll() ([]common.ProductResponse, error) {
	return []common.ProductResponse{
		{
			Id:       1,
			Merchant: "merchant1",
			Name:     "product1",
			Price:    100,
		},
		{
			Id:       2,
			Merchant: "merchant2",
			Name:     "product2",
			Price:    100,
		},
	}, nil
}

func (m mockProductRepositorySuccess) Get(int) (common.ProductResponse, error) {
	return common.ProductResponse{
		Id:       1,
		Merchant: "merchant1",
		Name:     "product1",
		Price:    100,
	}, nil
}

func (m mockProductRepositorySuccess) Create(entity.Product) (int, string, error) {
	return 1, "user1", nil
}

func (m mockProductRepositorySuccess) Update(entity.Product) (int, error) {
	return http.StatusOK, nil
}

func (m mockProductRepositorySuccess) Delete(int, int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllProductsSuccess(t *testing.T) {
	t.Run("TestGetAllProductsSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositorySuccess{})
		productController.GetAll()(context)

		actual := common.GetAllProductsResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllProductsResponse{
			Code:    http.StatusOK,
			Message: "get all products success",
			Data: []common.ProductResponse{
				{
					Id:       1,
					Merchant: "merchant1",
					Name:     "product1",
					Price:    100,
				},
				{
					Id:       2,
					Merchant: "merchant2",
					Name:     "product2",
					Price:    100,
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetProductSuccess(t *testing.T) {
	t.Run("TestGetProductSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositorySuccess{})
		productController.Get()(context)

		actual := common.GetProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetProductResponse{
			Code:    http.StatusOK,
			Message: "get product success",
			Data: []common.ProductResponse{
				{
					Id:       1,
					Merchant: "merchant1",
					Name:     "product1",
					Price:    100,
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateProductSuccess(t *testing.T) {
	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": 100,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositorySuccess{})
		midware.JWTMiddleware()(productController.Create())(context)

		actual := common.CreateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateProductResponse{
			Code:    http.StatusOK,
			Message: "create product success",
			Data: []common.ProductResponse{
				{
					Id:       1,
					Merchant: "user1",
					Name:     "product1",
					Price:    100,
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateProductSuccess(t *testing.T) {
	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": 100,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositorySuccess{})
		midware.JWTMiddleware()(productController.Update())(context)

		actual := common.UpdateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateProductResponse{
			Code:    http.StatusOK,
			Message: "update product success",
			Data: []entity.Product{
				{
					Id:     1,
					UserID: 1,
					Name:   "product1",
					Price:  100,
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteProductSuccess(t *testing.T) {
	t.Run("TestDeleteProductSuccess", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositorySuccess{})
		midware.JWTMiddleware()(productController.Delete())(context)

		actual := common.DeleteProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteProductResponse{
			Code:    http.StatusOK,
			Message: "delete product success",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN REPOSITORY

type mockProductRepositoryFailRepo struct{}

func (m mockProductRepositoryFailRepo) GetAll() ([]common.ProductResponse, error) {
	return nil, assert.AnError
}

func (m mockProductRepositoryFailRepo) Get(int) (common.ProductResponse, error) {
	return common.ProductResponse{}, assert.AnError
}

func (m mockProductRepositoryFailRepo) Create(entity.Product) (int, string, error) {
	return 0, "", assert.AnError
}

func (m mockProductRepositoryFailRepo) Update(entity.Product) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("update product failed")
}

func (m mockProductRepositoryFailRepo) Delete(int, int) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("delete product failed")
}

func TestGetAllProductsFailRepo(t *testing.T) {
	t.Run("TestGetAllProductsFailRepo", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositoryFailRepo{})
		productController.GetAll()(context)

		actual := common.GetAllProductsResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllProductsResponse{
			Code:    http.StatusInternalServerError,
			Message: "get all products failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetProductFailRepo(t *testing.T) {
	t.Run("TestGetProductFailRepo", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositoryFailRepo{})
		productController.Get()(context)

		actual := common.GetProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetProductResponse{
			Code:    http.StatusInternalServerError,
			Message: "get product failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateProductFailRepo(t *testing.T) {
	t.Run("TestCreateProductFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": 100,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositoryFailRepo{})
		midware.JWTMiddleware()(productController.Create())(context)

		actual := common.CreateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateProductResponse{
			Code:    http.StatusInternalServerError,
			Message: "create product failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateProductFailRepo(t *testing.T) {
	t.Run("TestUpdateProductFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": 100,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositoryFailRepo{})
		midware.JWTMiddleware()(productController.Update())(context)

		actual := common.UpdateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateProductResponse{
			Code:    http.StatusInternalServerError,
			Message: "update product failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteProductFailRepo(t *testing.T) {
	t.Run("TestDeleteProductFailRepo", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositoryFailRepo{})
		midware.JWTMiddleware()(productController.Delete())(context)

		actual := common.DeleteProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteProductResponse{
			Code:    http.StatusInternalServerError,
			Message: "delete product failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

// TEST FAIL IN BINDING, INVALID ID, OR EMPTY DIRECTORY

type mockProductRepositoryFailOther struct{}

func (m mockProductRepositoryFailOther) GetAll() ([]common.ProductResponse, error) {
	return []common.ProductResponse{}, nil
}

func (m mockProductRepositoryFailOther) Get(int) (common.ProductResponse, error) {
	return common.ProductResponse{}, nil
}

func (m mockProductRepositoryFailOther) Create(entity.Product) (int, string, error) {
	return 0, "", nil
}

func (m mockProductRepositoryFailOther) Update(entity.Product) (int, error) {
	return http.StatusOK, nil
}

func (m mockProductRepositoryFailOther) Delete(int, int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllProductsEmptyDirectory(t *testing.T) {
	t.Run("TestGetAllProductsEmptyDirectory", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositoryFailOther{})
		productController.GetAll()(context)

		actual := common.GetAllProductsResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetAllProductsResponse{
			Code:    http.StatusOK,
			Message: "products directory empty",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetProductFailInvalidId(t *testing.T) {
	t.Run("TestGetProductFailInvalidId", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		productController := New(mockProductRepositoryFailOther{})
		productController.Get()(context)

		actual := common.GetProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetProductResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid product id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetProductDoesNotExist(t *testing.T) {
	t.Run("TestGetProductDoesNotExist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositoryFailOther{})
		productController.Get()(context)

		actual := common.GetProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.GetProductResponse{
			Code:    http.StatusBadRequest,
			Message: "product does not exist",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateProductFailBinding(t *testing.T) {
	t.Run("TestCreateProductFailBinding", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": "100",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products")

		productController := New(mockProductRepositoryFailOther{})
		midware.JWTMiddleware()(productController.Create())(context)

		actual := common.CreateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.CreateProductResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateProductFailInvalidId(t *testing.T) {
	t.Run("TestUpdateProductFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": 100,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		productController := New(mockProductRepositoryFailRepo{})
		midware.JWTMiddleware()(productController.Update())(context)

		actual := common.UpdateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateProductResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid product id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateProductFailBinding(t *testing.T) {
	t.Run("TestUpdateProductFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "product1",
			"price": "1234",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := New(mockProductRepositoryFailRepo{})
		midware.JWTMiddleware()(productController.Update())(context)

		actual := common.UpdateProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.UpdateProductResponse{
			Code:    http.StatusBadRequest,
			Message: "binding failed",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteProductFailInvalidId(t *testing.T) {
	t.Run("TestDeleteProductFailInvalidId", func(t *testing.T) {
		token, _ := midware.CreateToken(1, "admin")

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("invalid id")

		productController := New(mockProductRepositoryFailOther{})
		midware.JWTMiddleware()(productController.Delete())(context)

		actual := common.DeleteProductResponse{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := common.DeleteProductResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid product id",
			Data:    nil,
		}

		assert.Equal(t, expected, actual)
	})
}
