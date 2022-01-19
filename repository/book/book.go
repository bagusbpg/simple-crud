package book

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api/design-pattern/delivery/common"
	"rest-api/design-pattern/entity"
)

type BookRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) GetAll() ([]common.BookResponse, error) {
	query := "SELECT id, title, author, publisher, language, pages, isbn13 FROM books"

	result, err := br.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	books := []common.BookResponse{}
	book := common.BookResponse{}

	for result.Next() {
		if err := result.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Language, &book.Pages, &book.ISBN13); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (br *BookRepository) Get(id int) (common.BookResponse, error) {
	book := common.BookResponse{}
	query := fmt.Sprintf("SELECT id, title, author, publisher, language, pages, isbn13 FROM books WHERE id=%v", id)

	result, err := br.db.Query(query)

	if err != nil {
		return book, err
	}

	defer result.Close()

	if !result.Next() {
		return book, err
	}

	if err := result.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Language, &book.Pages, &book.ISBN13); err != nil {
		return book, err
	}

	return book, nil
}

func (br *BookRepository) Create(book entity.Book) (int, error) {
	query := fmt.Sprintf("INSERT INTO books (title, author, publisher, language, pages, isbn13) VALUES ('%v','%v','%v','%v','%v','%v')", book.Title, book.Author, book.Publisher, book.Language, book.Pages, book.ISBN13)
	id := 0

	if _, err := br.db.Exec(query); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("SELECT id FROM books WHERE title='%v' AND author='%v' AND publisher='%v' AND language='%v' AND pages=%v AND isbn13='%v' ORDER BY id DESC LIMIT 1", book.Title, book.Author, book.Publisher, book.Language, book.Pages, book.ISBN13)

	result, _ := br.db.Query(query)
	defer result.Close()

	if result.Next() {
		result.Scan(&id)
	}

	return id, nil
}

func (br *BookRepository) Update(book entity.Book) (int, error) {
	query := fmt.Sprintf("UPDATE books SET title='%v', author='%v', publisher='%v', language='%v', pages=%v, isbn13='%v' WHERE id=%v", book.Title, book.Author, book.Publisher, book.Language, book.Pages, book.ISBN13, book.Id)

	result, err := br.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("udate book failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("udate book failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("book does not exist")
	}

	return http.StatusOK, nil
}

func (br *BookRepository) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM books WHERE id=%v", id)

	result, err := br.db.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete book failed")
	}

	count, err := result.RowsAffected()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("delete book failed")
	}

	if count == 0 {
		return http.StatusBadRequest, fmt.Errorf("book does not exist")
	}

	return http.StatusOK, nil
}
