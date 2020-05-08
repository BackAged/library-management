package author

import "context"

// Service provides port for application adapter.
type Service interface {
	Create(context.Context, *Author) error
	Get(context.Context, string) (*Author, error)
	List(context.Context, *int64, *int64) ([]Author, error)
}
