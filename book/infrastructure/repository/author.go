package repository

import (
	"context"
	"errors"
	"time"

	"github.com/BackAged/library-management/book/domain/author"
	"github.com/BackAged/library-management/book/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BsonBook defines bson book
type bsonAuthor struct {
	ID         primitive.ObjectID `bson:"_id"`
	AuthorName string             `bson:"author_name"`
	Details    string             `bson:"details"`
	CreatedAt  *time.Time         `bson:"created_at"`
	UpdatedAt  *time.Time         `bson:"updated_at"`
}

func (bBk *bsonAuthor) valid() (bool, error) {
	return true, nil
}

func toBsonAuthor(bk *author.Author) (*bsonAuthor, error) {
	bBk := bsonAuthor{
		AuthorName: bk.AuthorName,
		Details:    bk.Details,
		CreatedAt:  bk.CreatedAt,
		UpdatedAt:  bk.UpdatedAt,
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

func toModelAuthor(bk *bsonAuthor) *author.Author {
	return &author.Author{
		ID:         bk.ID.Hex(),
		AuthorName: bk.AuthorName,
		Details:    bk.Details,
		CreatedAt:  bk.CreatedAt,
		UpdatedAt:  bk.UpdatedAt,
	}
}

type authorRepository struct {
	db  *database.Client
	col string
}

// NewAuthorRepository returns a new AuthorRepository
func NewAuthorRepository(client *database.Client, col string) author.Repository {
	return &authorRepository{
		db:  client,
		col: col,
	}
}

// Add adds into repository
func (br *authorRepository) AddAuthor(ctx context.Context, bk *author.Author) error {
	bBk, err := toBsonAuthor(bk)
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

	*bk = *toModelAuthor(bBk)

	return nil
}

// Get fetches from repository
func (br *authorRepository) GetAuthor(ctx context.Context, ID string) (*author.Author, error) {
	if _, err := primitive.ObjectIDFromHex(ID); err != nil {
		return nil, nil
	}

	row, err := br.db.FindByID(context.Background(), br.col, ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	bBk := bsonAuthor{}

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

	return toModelAuthor(&bBk), nil
}

// Get fetches from repository
func (br *authorRepository) ListAuthor(ctx context.Context, skip *int64, limit *int64) ([]author.Author, error) {
	rows, err := br.db.Find(context.Background(), br.col, bson.M{}, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := []author.Author{}
	for rows.Next() {
		bBk := bsonAuthor{}
		if err := rows.Scan(&bBk); err != nil {
			return nil, err
		}
		bks = append(bks, *toModelAuthor(&bBk))

	}
	if rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
