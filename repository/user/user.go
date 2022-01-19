package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAll() ([]common.UserResponse, error) {
	query := "SELECT id, name, email FROM users"

	result, err := ur.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	users := []common.UserResponse{}
	user := common.UserResponse{}

	for result.Next() {
		if err := result.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) Get(id int) (common.UserResponse, error) {
	user := common.UserResponse{}
	query := fmt.Sprintf("SELECT id, name, email FROM users WHERE id=%v", id)

	result, err := ur.db.Query(query)

	if err != nil {
		return user, err
	}

	defer result.Close()

	if !result.Next() {
		return user, err
	}

	if err := result.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Create(user entity.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO users (name, email, password) VALUES ('%v','%v','%v')", user.Name, user.Email, user.Password)
	id := 0

	if _, err := ur.db.Exec(query); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("SELECT id FROM users WHERE name='%v' AND email='%v' AND password='%v' ORDER BY id DESC LIMIT 1", user.Name, user.Email, user.Password)

	result, _ := ur.db.Query(query)
	defer result.Close()

	if result.Next() {
		result.Scan(&id)
	}

	return id, nil
}

func (ur *UserRepository) Update(user entity.User) (int, error) {
	query := fmt.Sprintf("UPDATE users SET name='%v', email='%v', password='%v' WHERE id=%v", user.Name, user.Email, user.Password, user.Id)

	result, err := ur.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("update user failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("update user failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("user does not exist")
	}

	return http.StatusOK, nil
}

func (ur *UserRepository) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM users WHERE id=%v", id)

	result, err := ur.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete user failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete user failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("user does not exist")
	}

	return http.StatusOK, nil
}
