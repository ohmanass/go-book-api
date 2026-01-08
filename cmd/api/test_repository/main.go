//package testrepository
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/nassim-touissi/go-book-api/internal/model"
	"github.com/nassim-touissi/go-book-api/internal/repository"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresBookRepository(db)

	fmt.Println("Connected to DB")

	// Create
	year := 2020
	book := &model.Book{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Year:   &year,
	}

	if err := repo.Create(book); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book created:", book.ID)

	// GetByID
	found, err := repo.GetByID(book.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book fetched:", found.Title)

	// List
	books, err := repo.List(10, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Books count:", len(books))

	// Update
	book.Title = "Clean Code (Updated)"
	if err := repo.Update(book); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book updated")

	// Delete
	if err := repo.Delete(book.ID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book deleted")

	// Get after delete
	_, err = repo.GetByID(book.ID)
	if err == repository.ErrBookNotFound {
		fmt.Println("ErrBookNotFound works correctly")
	} else if err != nil {
		log.Fatal("Unexpected error:", err)
	}
}