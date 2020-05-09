package repository

// RepoValidationError represents repository validation error
type RepoValidationError struct {
	message string
	errors  []string
}

// NewRepoValidationError returns a new repository validation error
func NewRepoValidationError(message string, errors []string) *RepoValidationError {
	return &RepoValidationError{
		message: message,
		errors:  errors,
	}
}

func (rv *RepoValidationError) Error() string {
	return rv.message
}

func (rv *RepoValidationError) getError() []string {
	return rv.errors
}

func (rv *RepoValidationError) getMessage() string {
	return rv.message
}
