package author

import (
	"context"
)

type service struct {
	repository Repository
}

// NewService creates a service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Create(ctx context.Context, bk *Author) error {
	if err := s.repository.AddAuthor(ctx, bk); err != nil {
		return err
	}

	return nil
}

func (s *service) Get(ctx context.Context, ID string) (*Author, error) {
	bk, err := s.repository.GetAuthor(ctx, ID)
	if err != nil {
		return nil, err
	}
	if bk == nil {
		return nil, NewAuthorNotFound("Author not found", []string{})
	}

	return bk, nil
}

func (s *service) List(ctx context.Context, skip *int64, limit *int64) ([]Author, error) {
	bks, err := s.repository.ListAuthor(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return bks, nil
}
