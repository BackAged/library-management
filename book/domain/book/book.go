package book

import (
	"time"
)

// Book defines Book type
type Book struct {
	ID          string
	Title       string
	Category    string
	Description string
	AuthorID    string
	AuthorName  string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
