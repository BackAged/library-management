package bookloan

import "context"

// Service provides port for application adapter.
type Service interface {
	Create(context.Context, *BookLoan) error
	Get(context.Context, string) (*BookLoan, error)
	List(context.Context, *int64, *int64) ([]BookLoan, error)
	ListByUser(context.Context, string, *int64, *int64) ([]BookLoan, error)
	Accept(context.Context, string) error
	Reject(context.Context, string, string) error
}
