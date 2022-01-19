package book

import (
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type Book interface {
	GetAll() ([]common.BookResponse, error)
	Get(int) (common.BookResponse, error)
	Create(entity.Book) (int, error)
	Update(entity.Book) (int, error)
	Delete(int) (int, error)
}
