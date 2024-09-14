// internal/database/db.go

package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (db *DB) CreateBook(book *Book) error {
	return db.QueryRow(
		"INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id",
		book.Title, book.Author,
	).Scan(&book.ID)
}

func (db *DB) GetBook(id int) (*Book, error) {
	book := &Book{}
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (db *DB) UpdateBook(book *Book) error {
	_, err := db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3",
		book.Title, book.Author, book.ID)
	return err
}

func (db *DB) DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}

func (db *DB) ListBooks(limit, offset int) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author FROM books ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
