package product

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAll() ([]common.ProductResponse, error) {
	query := "SELECT p.id, u.name, p.name, p.price FROM products p LEFT JOIN users u ON p.user_id = u.id"

	result, err := pr.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	products := []common.ProductResponse{}
	product := common.ProductResponse{}

	for result.Next() {
		if err := result.Scan(&product.Id, &product.Merchant, &product.Name, &product.Price); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) Get(id int) (common.ProductResponse, error) {
	product := common.ProductResponse{}
	query := fmt.Sprintf("SELECT p.id, u.name, p.name, p.price FROM products p LEFT JOIN users u ON p.user_id = u.id WHERE p.id=%v", id)

	result, err := pr.db.Query(query)

	if err != nil {
		return product, err
	}

	defer result.Close()

	if !result.Next() {
		return product, err
	}

	if err := result.Scan(&product.Id, &product.Merchant, &product.Name, &product.Price); err != nil {
		return product, err
	}

	return product, nil
}

func (pr *ProductRepository) Create(product entity.Product) (int, string, error) {
	query := fmt.Sprintf("INSERT INTO products (user_id, name, price) VALUES ('%v', '%v', '%v')", product.UserID, product.Name, product.Price)
	id := 0

	if _, err := pr.db.Exec(query); err != nil {
		return id, "", err
	}

	query = fmt.Sprintf("SELECT id FROM products WHERE name='%v' AND user_id=%v AND price='%v' ORDER BY id DESC LIMIT 1", product.Name, product.UserID, product.Price)

	result, _ := pr.db.Query(query)
	defer result.Close()

	if result.Next() {
		result.Scan(&id)
	}

	query = fmt.Sprintf("SELECT name FROM users WHERE id=%v", product.UserID)

	result, _ = pr.db.Query(query)
	defer result.Close()

	name := ""

	if result.Next() {
		result.Scan(&name)
	}

	return id, name, nil
}

func (pr *ProductRepository) Update(product entity.Product) (int, error) {
	query := fmt.Sprintf("UPDATE products SET name='%v', price='%v' WHERE id=%v AND user_id=%v", product.Name, product.Price, product.Id, product.UserID)

	result, err := pr.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("update product failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("update product failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("product does not exist")
	}

	return http.StatusOK, nil
}

func (pr *ProductRepository) Delete(id int, userid int) (int, error) {
	query := fmt.Sprintf("DELETE FROM products WHERE id=%v AND user_id=%v", id, userid)

	result, err := pr.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete product failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete product failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("user does not match")
	}

	return http.StatusOK, nil
}
