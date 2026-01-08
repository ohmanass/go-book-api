package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL connection parameters
	const (
		host     = "localhost"
		port     = 5434
		user     = "admin"
		password = "admin"
		dbname   = "booksdb"
	)

	// Build connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Open database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close connection on exit

	// Verify database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Initialize HTTP router
	r := chi.NewRouter()

	// Health check route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Book CRUD routes (handlers to be implemented)
	r.Post("/books", createBookHandler)
	r.Get("/books", listBookHandler)
	r.Get("/books/{id}", getBookHandler)
	r.Put("/books/{id}", updateBookHandler)
	r.Delete("/books/{id}", deleteBookHandler)

	// Start HTTP server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// Empty handlers for compilation
func createBookHandler(w http.ResponseWriter, r *http.Request) {}
func listBookHandler(w http.ResponseWriter, r *http.Request)   {}
func getBookHandler(w http.ResponseWriter, r *http.Request)    {}
func updateBookHandler(w http.ResponseWriter, r *http.Request) {}
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {}