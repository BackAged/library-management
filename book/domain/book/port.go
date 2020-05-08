package book

import "context"

// Repository defines the port of book for infrastracture adapter
type Repository interface {
	AddBook(context.Context, *Book) error
	GetBook(context.Context, string) (*Book, error)
	ListBook(context.Context, *int64, *int64) ([]Book, error)
	ListBookByAutherID(context.Context, string, *int64, *int64) ([]Book, error)
}
