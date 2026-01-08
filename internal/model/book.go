package model

import (
	"errors"
	"time"
	"strings"
	"fmt"
)

type Book struct {
	ID        string    `json:"id"`        
	Title     string    `json:"title"`   
	Author    string    `json:"author"`    
	Year      *int      `json:"year"`      
	CreatedAt time.Time `json:"createdAt"` 
	UpdatedAt time.Time `json:"updatedAt"` 
}

type CreateBookRequest struct {
	Title  string `json:"title"`           
	Author string `json:"author"`        
	Year   *int   `json:"year,omitempty"` 
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   *int   `json:"year,omitempty"`
}

func (b *Book) Validate() error {
    if strings.TrimSpace(b.Title) == "" {
		return errors.New("Un titre est requis !")
	}

	if strings.TrimSpace(b.Author) == "" {
		return errors.New("Un autheur est requis !")
	}

	if b.Year != nil {
		currentDate := time.Now().Year()
		if *b.Year < 700 || *b.Year > currentDate {
			return fmt.Errorf("Le livre doit être paru entre 700 et %d", currentDate)
		} 
	}
	return nil
}