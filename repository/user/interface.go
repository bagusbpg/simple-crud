package user

import (
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type User interface {
	GetAll() ([]common.UserResponse, error)
	Get(int) (common.UserResponse, error)
	Create(entity.User) (int, error)
	Update(entity.User) (int, error)
	Delete(int) (int, error)
}
