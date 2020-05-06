package book

import "context"

// Service provides port for application adapter.
type Service interface {
	Create(context.Context, *Book) error
	// Get(context.Context, string) (*Book, error)
	// GetAuthorBooks(context.Context, string) ([]*Book, error)
}
