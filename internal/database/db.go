// internal/database/db.go

package database

import (
	"context"
	"database/sql"
	"time"
	"fmt"

	"skylab-gin-book-api/internal/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
	Redis *redis.Client
}

func NewDB(dbConnString, cacheConnString string) (*DB, error) {
    cfg, err := config.LoadConfig()
	
	fmt.Println("Attempting to connect to postgres...")
	db, err := sql.Open("postgres", dbConnString)
    if err != nil {
		fmt.Errorf("error accessing postgres instance: %w", err)
        return nil, err
    } else {
		fmt.Println("Successfully connected to the postgres instance.")
	}
    if err = db.Ping(); err != nil {
		fmt.Println("error accessing postgres instance: %w", err)
        return nil, err
    }

	var rdb *redis.Client
	if cfg.Cache.UseCache {
		fmt.Println("Attempting to connect to redis...")
		rdb := redis.NewClient(&redis.Options{
			Addr: cacheConnString,
		})
		
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		_, err = rdb.Ping(ctx).Result()
		if err != nil {
		fmt.Println("Warning: Redis is not available. The application will continue without caching.")
		} else {
			fmt.Println("Successfully connected to Redis.")
		}
	} else {
		fmt.Println("Redis is disabled in the configuration. The application will run without caching.")
	}

    return &DB{DB: db, Redis: rdb}, nil
}

func CreateTable(db *DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL
    );
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	fmt.Println("Postgres tables staged or already present.")
	return nil
}

type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
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

// Redis helper functions

func (db *DB) SetCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return db.Redis.Set(ctx, key, value, expiration).Err()
}

func (db *DB) GetCache(key string) (string, error) {
	ctx := context.Background()
	return db.Redis.Get(ctx, key).Result()
}

func (db *DB) DeleteCache(key string) error {
	ctx := context.Background()
	return db.Redis.Del(ctx, key).Err()
}
