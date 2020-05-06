package book

import "context"

// Repository defines the port of book for infrastracture adapter
type Repository interface {
	Add(context.Context, *Book) error
	// Get(context.Context, string) (*Book, error)
	// FindByAutherID(context.Context, string) ([]*Book, error)
}
