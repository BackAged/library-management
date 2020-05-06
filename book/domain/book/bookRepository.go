package book

import "context"

// Repository defines the port of book for infrastracture adapter
type Repository interface {
	Add(context.Context, *Book) error
	FindByID(context.Context, string) (*Book, error)
	FindByUserID(context.Context, string) ([]*Book, error)
}
