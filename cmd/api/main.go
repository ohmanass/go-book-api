package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"github.com/nassim-touissi/go-book-api/internal/handler"
	"github.com/nassim-touissi/go-book-api/internal/repository"
)

func main() {
	// Build PostgreSQL connection string from environment variables
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Open database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	// ============================= Test query: fetch one book
	var id, title, author string
	var year sql.NullInt32

	err = db.QueryRow("SELECT id, title, author, year FROM books LIMIT 1").Scan(&id, &title, &author, &year)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No books found")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Book example: ID=%s, Title=%s, Author=%s, Year=%v\n", id, title, author, year)
	}

	// ============================= Initialize repository and handlers
	repo := repository.NewPostgresBookRepository(db)
	bookHandler := handler.NewBookHandler(repo)

	// Initialize HTTP router
	r := chi.NewRouter()

	// Book CRUD routes
	r.Post("/books", bookHandler.CreateBookHandler)
	r.Get("/books", bookHandler.ListBooksHandler)
	r.Get("/books/{id}", bookHandler.GetBookHandler)
	r.Put("/books/{id}", bookHandler.UpdateBookHandler)
	r.Delete("/books/{id}", bookHandler.DeleteBookHandler)

	// Health check route
	r.Get("/health", handler.HealthHandler(db))

	// ============================= Start HTTP server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}