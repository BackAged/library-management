package bookloan

import "context"

// Repository defines the port of book for infrastracture adapter
type Repository interface {
	AddBookLoan(context.Context, *BookLoan) error
	GetBookLoan(context.Context, string) (*BookLoan, error)
	ListBookLoan(context.Context, *int64, *int64) ([]BookLoan, error)
	ListBookLoanByUserID(context.Context, string, *int64, *int64) ([]BookLoan, error)
	UpdateBookLoan(context.Context, string, *BookLoan) error
}
