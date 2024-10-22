package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"skylab-gin-book-api/internal/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

const (
	defaultCacheExpiration = 24 * time.Hour
	bookCacheKeyPrefix    = "book:"
	booksListCacheKey     = "books:list:"
)

type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

type DB struct {
	*sql.DB
	Redis     *redis.Client
	UseCache  bool
}

func NewDB(dbConnString, cacheConnString string) (*DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	fmt.Println("Attempting to connect to postgres...")
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, fmt.Errorf("error accessing postgres instance: %w", err)
	}
	
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging postgres instance: %w", err)
	}
	fmt.Println("Successfully connected to the postgres instance.")

	var rdb *redis.Client
	useCache := cfg.Cache.UseCache

	if useCache {
		fmt.Println("Attempting to connect to redis...")
		rdb = redis.NewClient(&redis.Options{
			Addr: cacheConnString,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = rdb.Ping(ctx).Result()
		if err != nil {
			fmt.Println("Warning: Redis is not available. The application will continue without caching.")
			useCache = false
		} else {
			fmt.Println("Successfully connected to Redis.")
		}
	} else {
		fmt.Println("Redis is disabled in the configuration. The application will run without caching.")
	}

	return &DB{
		DB:       db,
		Redis:    rdb,
		UseCache: useCache,
	}, nil
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

func (db *DB) CreateBook(book *Book) error {
	err := db.QueryRow(
		"INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id",
		book.Title, book.Author,
	).Scan(&book.ID)

	if err != nil {
		return err
	}

	// Invalidate the list cache when a new book is added
	if db.UseCache {
		ctx := context.Background()
		db.Redis.Del(ctx, booksListCacheKey+"*")
	}

	return nil
}

func (db *DB) GetBook(id int) (*Book, error) {
	if db.UseCache {
		// Try to get from cache first
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d", bookCacheKeyPrefix, id)
		
		cachedBook, err := db.Redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var book Book
			if err := json.Unmarshal([]byte(cachedBook), &book); err == nil {
				return &book, nil
			}
		}
	}

	// If not in cache or cache disabled, get from database
	book := &Book{}
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}

	// Store in cache if enabled
	if db.UseCache {
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d", bookCacheKeyPrefix, id)
		
		if bookJSON, err := json.Marshal(book); err == nil {
			db.Redis.Set(ctx, cacheKey, bookJSON, defaultCacheExpiration)
		}
	}

	return book, nil
}

func (db *DB) UpdateBook(book *Book) error {
	_, err := db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3",
		book.Title, book.Author, book.ID)
	
	if err != nil {
		return err
	}

	// Invalidate cache if enabled
	if db.UseCache {
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d", bookCacheKeyPrefix, book.ID)
		db.Redis.Del(ctx, cacheKey)
		// Invalidate list cache as well
		db.Redis.Del(ctx, booksListCacheKey+"*")
	}

	return nil
}

func (db *DB) DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	
	if err != nil {
		return err
	}

	// Invalidate cache if enabled
	if db.UseCache {
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d", bookCacheKeyPrefix, id)
		db.Redis.Del(ctx, cacheKey)
		// Invalidate list cache as well
		db.Redis.Del(ctx, booksListCacheKey+"*")
	}

	return nil
}

func (db *DB) ListBooks(limit, offset int) ([]Book, error) {
	if db.UseCache {
		// Try to get from cache first
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d:%d", booksListCacheKey, limit, offset)
		
		cachedBooks, err := db.Redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var books []Book
			if err := json.Unmarshal([]byte(cachedBooks), &books); err == nil {
				return books, nil
			}
		}
	}

	// If not in cache or cache disabled, get from database
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

	// Store in cache if enabled
	if db.UseCache {
		ctx := context.Background()
		cacheKey := fmt.Sprintf("%s%d:%d", booksListCacheKey, limit, offset)
		
		if booksJSON, err := json.Marshal(books); err == nil {
			db.Redis.Set(ctx, cacheKey, booksJSON, defaultCacheExpiration)
		}
	}

	return books, nil
}