package bookloan

import (
	"time"
)

// LoanStatus defines LoanStatus type
type LoanStatus string

// LoanStatus enums
const (
	BookLoanInitiated LoanStatus = "Initiated"
	BookLoanPending   LoanStatus = "Pending"
	BookLoanAccepted  LoanStatus = "Accepted"
	BookLoanRejected  LoanStatus = "Rejected"
)

// BookLoan defines BookLoan type
type BookLoan struct {
	ID             string
	BookID         string
	UserID         string
	Status         LoanStatus
	AcceptedAt     *time.Time
	RejectedAt     *time.Time
	RejectionCause string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func (bl *BookLoan) valid() (bool, error) {
	return true, nil
}
