package book

import (
	"context"
	"fmt"
)

type service struct {
	repository Repository
}

// NewService creates a service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Create(ctx context.Context, bk *Book) error {
	fmt.Print(bk)
	if err := s.repository.Add(ctx, bk); err != nil {
		return err
	}

	return nil
}

// func (s *service) Get(ctx context.Context, ID string) (*Book, error) {
// 	tsk, err := s.repository.FindByID(ctx, ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tsk, nil
// }

// func (s *service) GetUserTask(ctx context.Context, userID string) ([]*Book, error) {
// 	tsks, err := s.repository.FindByUserID(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tsks, nil
// }
