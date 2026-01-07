package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"log"
)

func main() {
	r := chi.NewRouter() // router creation

	// Route health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Routes books
	r.Post("/books", createBookHandler)
	r.Get("/books", listBooksHandler)
	r.Get("/books/{id}", getBookHandler)
	r.Put("/books/{id}", updateBookHandler)
	r.Delete("/books/{id}", deleteBookHandler)

	// Démarrage du serveur
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

// Handlers empty for compilation
func createBookHandler(w http.ResponseWriter, r *http.Request) {}
func listBooksHandler(w http.ResponseWriter, r *http.Request) {}
func getBookHandler(w http.ResponseWriter, r *http.Request) {}
func updateBookHandler(w http.ResponseWriter, r *http.Request) {}
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {}