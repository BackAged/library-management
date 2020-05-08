package repository

import (
	"context"
	"errors"
	"time"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/BackAged/library-management/book/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BsonBook defines bson book
type bsonBook struct {
	ID          primitive.ObjectID `bson:"_id"`
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

func toBosonBook(bk *book.Book) (*bsonBook, error) {
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

func toModelBook(bk *bsonBook) *book.Book {
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
func (br *bookRepository) AddBook(ctx context.Context, bk *book.Book) error {
	bBk, err := toBosonBook(bk)
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

	*bk = *toModelBook(bBk)

	return nil
}

// Get fetches from repository
func (br *bookRepository) GetBook(ctx context.Context, ID string) (*book.Book, error) {
	if _, err := primitive.ObjectIDFromHex(ID); err != nil {
		return nil, nil
	}

	row, err := br.db.FindByID(context.Background(), br.col, ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	bBk := bsonBook{}

	if row.Next() {
		if err := row.Scan(&bBk); err != nil {
			return nil, err
		}
	}

	if err := row.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return toModelBook(&bBk), nil
}

// Get fetches from repository
func (br *bookRepository) ListBook(ctx context.Context, skip *int64, limit *int64) ([]book.Book, error) {
	rows, err := br.db.Find(context.Background(), br.col, bson.M{}, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := []book.Book{}
	for rows.Next() {
		bBk := bsonBook{}
		if err := rows.Scan(&bBk); err != nil {
			return nil, err
		}
		bks = append(bks, *toModelBook(&bBk))

	}
	if rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

func (br *bookRepository) ListBookByAutherID(ctx context.Context, authorID string, skip *int64, limit *int64) ([]book.Book, error) {
	query := bson.M{
		"author_id": authorID,
	}

	rows, err := br.db.Find(context.Background(), br.col, query, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := []book.Book{}
	for rows.Next() {
		bBk := bsonBook{}
		if err := rows.Scan(&bBk); err != nil {
			return nil, err
		}
		bks = append(bks, *toModelBook(&bBk))

	}
	if rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
