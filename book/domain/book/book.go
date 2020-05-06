package book

import (
	"errors"
	"time"
)

// Lit bit of ddd modeling

// Status defines Book valid status
type Status string

// Status all type
const (
	Pending  Status = "Pending"
	Active   Status = "Active"
	InActive Status = "InActive"
	Complete Status = "Complete"
)

// Book defines Book type
type Book struct {
	ID          string
	UserID      string
	Topic       string
	Description string
	Status      Status
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// Complete makes the Book status Complete
func (t *Book) Complete() error {
	if t.Status == Complete || t.Status == InActive {
		return errors.New("Can't complete a inActvie Or complete Book")
	}

	t.Status = Complete
	return nil
}

