package author

import "encoding/json"

// NotFound service error
type NotFound struct {
	message string
	errors  []string
}

func (anf *NotFound) Error() string {
	if len(anf.errors) != 0 {
		b, _ := json.Marshal(anf.errors)
		return string(b)
	}

	return anf.message
}

// Add adds new error
func (anf *NotFound) Add(key string, value string) {
	anf.errors = append(anf.errors, value)
}

// GetMessage returns error message
func (anf *NotFound) GetMessage() string {
	return anf.message
}

// GetErrors returns erros
func (anf *NotFound) GetErrors() []string {
	return anf.errors
}

// NewBookNotFound returns new BookNotFound error
func NewAuthorNotFound(message string, errors []string) *NotFound {
	return &NotFound{
		message: message,
		errors:  errors,
	}
}
