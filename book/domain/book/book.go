package book

import (
	"time"
)

// Book defines Book type
type Book struct {
	ID          string
	Title       string
	ISBN        string
	Category    string
	Description string
	AuthorID    string
	AuthorName  string
	Quantity    int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (bk *Book) valid() (bool, error) {
	return true, nil
}
