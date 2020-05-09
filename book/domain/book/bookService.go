package book

import (
	"context"

	"github.com/BackAged/library-management/book/domain/author"
)

type service struct {
	repository Repository
	athrRepo   author.Repository
}

// NewService creates a service with the necessary dependencies.
func NewService(r Repository, authrRepo author.Repository) Service {
	return &service{
		repository: r,
		athrRepo:   authrRepo,
	}
}

func (s *service) Create(ctx context.Context, bk *Book) error {
	athr, err := s.athrRepo.GetAuthor(ctx, bk.AuthorID)
	if err != nil {
		return err
	}
	if athr == nil {
		return NewAuthorNotFound("Invalid author", []string{})
	}

	// TODO:-> check ISBN uniqueness

	bk.AuthorName = athr.AuthorName

	if err := s.repository.AddBook(ctx, bk); err != nil {
		return err
	}

	return nil
}

func (s *service) Get(ctx context.Context, ID string) (*Book, error) {
	bk, err := s.repository.GetBook(ctx, ID)
	if err != nil {
		return nil, err
	}
	if bk == nil {
		return nil, NewBookNotFound("Book not found", []string{})
	}

	return bk, nil
}

func (s *service) List(ctx context.Context, skip *int64, limit *int64) ([]Book, error) {
	bks, err := s.repository.ListBook(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

func (s *service) GetAuthorBooks(ctx context.Context, authorID string, skip *int64, limit *int64) ([]Book, error) {
	bks, err := s.repository.ListBookByAutherID(ctx, authorID, skip, limit)
	if err != nil {
		return nil, err
	}

	return bks, nil
}
