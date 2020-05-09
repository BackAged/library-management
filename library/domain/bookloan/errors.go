package bookloan

import "encoding/json"

// NotFound service error
type NotFound struct {
	message string
	errors  []string
}

func (bnf *NotFound) Error() string {
	if len(bnf.errors) != 0 {
		b, _ := json.Marshal(bnf.errors)
		return string(b)
	}

	return bnf.message
}

// Add adds new error
func (bnf *NotFound) Add(key string, value string) {
	bnf.errors = append(bnf.errors, value)
}

// GetMessage returns error message
func (bnf *NotFound) GetMessage() string {
	return bnf.message
}

// GetErrors returns erros
func (bnf *NotFound) GetErrors() []string {
	return bnf.errors
}

// NewBookLoanNotFound returns new BookLoanNotFound error
func NewBookLoanNotFound(message string, errors []string) *NotFound {
	return &NotFound{
		message: message,
		errors:  errors,
	}
}
