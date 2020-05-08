package author

import "context"

// Repository defines the port of book for infrastracture adapter
type Repository interface {
	AddAuthor(context.Context, *Author) error
	GetAuthor(context.Context, string) (*Author, error)
	ListAuthor(context.Context, *int64, *int64) ([]Author, error)
}
