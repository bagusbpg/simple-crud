package product

import (
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/delivery/midware"
	"rest-api/design-pattern/entity"
	productRepo "rest-api/design-pattern/repository/product"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	repository productRepo.Product
}

func New(product productRepo.Product) *ProductController {
	return &ProductController{
		repository: product,
	}
}

func (pc ProductController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := pc.repository.GetAll()
		code := http.StatusOK

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get all products failed", nil))
		}

		if len(products) == 0 {
			return c.JSON(code, common.SimpleResponse(code, "products directory empty", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get all products success", products))
	}
}

func (pc ProductController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		code := http.StatusOK

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid product id", nil))
		}

		product, err := pc.repository.Get(id)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "get product failed", nil))
		}

		if product == (common.ProductResponse{}) {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "product does not exist", nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "get product success", []common.ProductResponse{product}))
	}
}

func (pc ProductController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := midware.ExtractId(c)
		code := http.StatusOK

		if err != nil {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		input := entity.Product{}
		input.UserID = userid

		if err := c.Bind(&input); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		id, name, err := pc.repository.Create(input)

		if err != nil {
			code = http.StatusInternalServerError
			return c.JSON(code, common.SimpleResponse(code, "create product failed", nil))
		}

		product := common.ProductResponse{}
		product.Id = id
		product.Merchant = name
		product.Name = input.Name
		product.Price = input.Price

		return c.JSON(code, common.SimpleResponse(code, "create product success", []common.ProductResponse{product}))
	}
}

func (pc ProductController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := midware.ExtractId(c)
		code := http.StatusOK

		if err != nil {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		product := entity.Product{}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid product id", nil))
		}

		if err := c.Bind(&product); err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "binding failed", nil))
		}

		product.UserID = userid
		product.Id = id

		if code, err := pc.repository.Update(product); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "update product success", []entity.Product{product}))
	}
}

func (pc ProductController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := midware.ExtractId(c)
		code := http.StatusOK

		if err != nil {
			code = http.StatusUnauthorized
			return c.JSON(code, common.SimpleResponse(code, "unauthorized", nil))
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			code = http.StatusBadRequest
			return c.JSON(code, common.SimpleResponse(code, "invalid product id", nil))
		}

		if code, err := pc.repository.Delete(id, userid); err != nil {
			return c.JSON(code, common.SimpleResponse(code, err.Error(), nil))
		}

		return c.JSON(code, common.SimpleResponse(code, "delete product success", nil))
	}
}
