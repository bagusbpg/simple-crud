package product

import (
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type Product interface {
	GetAll() ([]common.ProductResponse, error)
	Get(int) (common.ProductResponse, error)
	Create(entity.Product) (int, string, error)
	Update(entity.Product) (int, error)
	Delete(int, int) (int, error)
}
