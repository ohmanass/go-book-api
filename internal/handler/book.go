package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nassim-touissi/go-book-api/internal/model"
	"github.com/nassim-touissi/go-book-api/internal/repository"
)

// BookHandler holds a reference to the repository
type BookHandler struct {
	repo repository.BookRepository
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

// POST /books - create a new book
func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	book := &model.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := book.Validate(); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(book); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GET /books - list books with optional pagination
func (h *BookHandler) ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Optional query parameters
	limit := 10
	offset := 0

	books, err := h.repo.List(limit, offset)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET /books/{id} - get a book by ID
func (h *BookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := h.repo.GetByID(id)
	if err != nil {
		if err == repository.ErrBookNotFound {
			http.Error(w, `{"error":"Book not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// PUT /books/{id} - update a book completely
func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req model.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	book := &model.Book{
		ID:     id,
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := book.Validate(); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(book); err != nil {
		if err == repository.ErrBookNotFound {
			http.Error(w, `{"error":"Book not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// DELETE /books/{id} - delete a book by ID
func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.repo.Delete(id); err != nil {
		if err == repository.ErrBookNotFound {
			http.Error(w, `{"error":"Book not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}