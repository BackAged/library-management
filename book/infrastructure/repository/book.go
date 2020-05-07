package repository

import (
	"context"
	"errors"
	"time"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/BackAged/library-management/book/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BsonBook defines bson book
type bsonBook struct {
	ID          primitive.ObjectID `bson:"id"`
	Title       string             `bson:"title"`
	ISBN        string             `bson:"isbn"`
	Category    string             `bson:"category"`
	Description string             `bson:"description"`
	AuthorID    string             `bson:"author_id"`
	AuthorName  string             `bson:"author_name"`
	Quantity    int                `bson:"quantity"`
	CreatedAt   *time.Time         `bson:"created_at"`
	UpdatedAt   *time.Time         `bson:"updated_at"`
}

func (bBk *bsonBook) valid() (bool, error) {
	return true, nil
}

func toBson(bk *book.Book) (*bsonBook, error) {
	bBk := bsonBook{
		Title:       bk.Title,
		ISBN:        bk.ISBN,
		Category:    bk.Category,
		Description: bk.Description,
		AuthorID:    bk.AuthorID,
		AuthorName:  bk.AuthorName,
		Quantity:    bk.Quantity,
		CreatedAt:   bk.CreatedAt,
		UpdatedAt:   bk.UpdatedAt,
	}

	if bk.ID != "" {
		id, err := primitive.ObjectIDFromHex(bk.ID)
		if err != nil {
			return nil, errors.New("invalid id")
		}
		bBk.ID = id
	}
	return &bBk, nil
}

func toModel(bk *bsonBook) *book.Book {
	return &book.Book{
		ID:          bk.ID.Hex(),
		Title:       bk.Title,
		ISBN:        bk.ISBN,
		Category:    bk.Category,
		Description: bk.Description,
		AuthorID:    bk.AuthorID,
		AuthorName:  bk.AuthorName,
		Quantity:    bk.Quantity,
		CreatedAt:   bk.CreatedAt,
		UpdatedAt:   bk.UpdatedAt,
	}
}

type bookRepository struct {
	db  *database.Client
	col string
}

// NewBookRepository returns a new taskRepository
func NewBookRepository(client *database.Client, col string) book.Repository {
	return &bookRepository{
		db:  client,
		col: col,
	}
}

// Add adds into repository
func (br *bookRepository) Add(ctx context.Context, bk *book.Book) error {
	bBk, err := toBson(bk)
	if err != nil {
		return err
	}

	now := time.Now()

	if bBk.ID.IsZero() {
		bBk.ID = primitive.NewObjectID()
	}
	bBk.CreatedAt = &now
	bBk.UpdatedAt = &now

	if ok, err := bBk.valid(); !ok {
		return err
	}

	if _, err := br.db.Insert(context.Background(), br.col, bBk); err != nil {
		return err
	}

	*bk = *toModel(bBk)

	return nil
}
