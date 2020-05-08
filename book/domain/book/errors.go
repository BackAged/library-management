package book

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

// NewBookNotFound returns new BookNotFound error
func NewBookNotFound(message string, errors []string) *NotFound {
	return &NotFound{
		message: message,
		errors:  errors,
	}
}

// AuthorNotFound service error
type AuthorNotFound struct {
	message string
	errors  []string
}

func (anf *AuthorNotFound) Error() string {
	if len(anf.errors) != 0 {
		b, _ := json.Marshal(anf.errors)
		return string(b)
	}

	return anf.message
}

// Add adds new error
func (anf *AuthorNotFound) Add(key string, value string) {
	anf.errors = append(anf.errors, value)
}

// GetMessage returns error message
func (anf *AuthorNotFound) GetMessage() string {
	return anf.message
}

// GetErrors returns erros
func (anf *AuthorNotFound) GetErrors() []string {
	return anf.errors
}

// NewAuthorNotFound returns new BookNotFound error
func NewAuthorNotFound(message string, errors []string) *AuthorNotFound {
	return &AuthorNotFound{
		message: message,
		errors:  errors,
	}
}
