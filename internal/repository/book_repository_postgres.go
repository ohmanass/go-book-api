package repository

import (
	"context"
	"database/sql"
	"time"
	"github.com/google/uuid"
	"github.com/nassim-touissi/go-book-api/internal/model"
)

// PostgresBookRepository implements BookRepository using PostgreSQL
type PostgresBookRepository struct {
	db *sql.DB
}

// Creates a new PostgreSQL-backed BookRepository
func NewPostgresBookRepository(db *sql.DB) BookRepository {
	return &PostgresBookRepository{db: db}
}

// Inserts a new book into the database

func (r *PostgresBookRepository) Create(book *model.Book) error {
	if book.ID == "" {
		book.ID = uuid.New().String() // generate UUID if empty
	}

	now := time.Now()
	book.CreatedAt = now
	book.UpdatedAt = now

	query := `
		INSERT INTO books (id, title, author, year, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		context.Background(),
		query,
		book.ID,
		book.Title,
		book.Author,
		book.Year,
		book.CreatedAt,
		book.UpdatedAt,
	)

	return err
}

// Retrieves a book by its ID
func (r *PostgresBookRepository) GetByID(id string) (*model.Book, error) {
	query := `
		SELECT id, title, author, year, created_at, updated_at
		FROM books
		WHERE id = $1
	`

	var book model.Book

	err := r.db.QueryRowContext(context.Background(), query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Year,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return &book, nil
}

// Returns a list of books with pagination support
func (r *PostgresBookRepository) List(limit, offset int) ([]*model.Book, error) {
	query := `
		SELECT id, title, author, year, created_at, updated_at
		FROM books
		ORDER BY created_at
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*model.Book

	for rows.Next() {
		var book model.Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// Updates an existing book
func (r *PostgresBookRepository) Update(book *model.Book) error {
	query := `
		UPDATE books
		SET title = $1, author = $2, year = $3, updated_at = $4
		WHERE id = $5
	`

	book.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(
		context.Background(),
		query,
		book.Title,
		book.Author,
		book.Year,
		book.UpdatedAt,
		book.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}

// Deletes a book by its ID
func (r *PostgresBookRepository) Delete(id string) error {
	query := `DELETE FROM books WHERE id = $1`

	result, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}