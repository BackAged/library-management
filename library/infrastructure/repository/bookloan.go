package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BackAged/library-management/library/domain/bookloan"
	"github.com/BackAged/library-management/library/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// bsonBookLoan defines bson book
type bsonBookLoan struct {
	ID             primitive.ObjectID `bson:"_id"`
	BookID         string             `bson:"book_id"`
	UserID         string             `bson:"user_id"`
	Status         string             `bson:"status"`
	AcceptedAt     *time.Time         `bson:"accepted_at"`
	RejectedAt     *time.Time         `bson:"rejected_at"`
	RejectionCause string             `bson:"rejection_cause"`
	CreatedAt      *time.Time         `bson:"created_at"`
	UpdatedAt      *time.Time         `bson:"updated_at"`
}

func (bl *bsonBookLoan) valid() (bool, error) {
	return true, nil
}

func toBsonBookLoan(bl *bookloan.BookLoan) (*bsonBookLoan, error) {
	bBl := bsonBookLoan{
		BookID:         bl.BookID,
		UserID:         bl.UserID,
		Status:         string(bl.Status),
		AcceptedAt:     bl.AcceptedAt,
		RejectedAt:     bl.RejectedAt,
		RejectionCause: bl.RejectionCause,
		CreatedAt:      bl.CreatedAt,
		UpdatedAt:      bl.UpdatedAt,
	}

	if bl.ID != "" {
		id, err := primitive.ObjectIDFromHex(bl.ID)
		if err != nil {
			return nil, errors.New("invalid id")
		}
		bBl.ID = id
	}
	return &bBl, nil
}

func toModelBookLoan(bl *bsonBookLoan) *bookloan.BookLoan {
	return &bookloan.BookLoan{
		ID:             bl.ID.Hex(),
		BookID:         bl.BookID,
		UserID:         bl.UserID,
		Status:         bookloan.LoanStatus(bl.Status),
		AcceptedAt:     bl.AcceptedAt,
		RejectedAt:     bl.RejectedAt,
		RejectionCause: bl.RejectionCause,
		CreatedAt:      bl.CreatedAt,
		UpdatedAt:      bl.UpdatedAt,
	}
}

type bookLoanRepository struct {
	db  *database.Client
	col string
}

// NewBookLoanRepository returns a new AuthorRepository
func NewBookLoanRepository(client *database.Client, col string) bookloan.Repository {
	return &bookLoanRepository{
		db:  client,
		col: col,
	}
}

// AddBookLoan adds into repository
func (br *bookLoanRepository) AddBookLoan(ctx context.Context, bk *bookloan.BookLoan) error {
	bBl, err := toBsonBookLoan(bk)
	if err != nil {
		return err
	}

	now := time.Now()

	if bBl.ID.IsZero() {
		bBl.ID = primitive.NewObjectID()
	}
	bBl.CreatedAt = &now
	bBl.UpdatedAt = &now

	if ok, err := bBl.valid(); !ok {
		return err
	}

	if _, err := br.db.Insert(context.Background(), br.col, bBl); err != nil {
		return err
	}

	*bk = *toModelBookLoan(bBl)

	return nil
}

// GetBookLoan fetches from repository
func (br *bookLoanRepository) GetBookLoan(ctx context.Context, ID string) (*bookloan.BookLoan, error) {
	if _, err := primitive.ObjectIDFromHex(ID); err != nil {
		return nil, nil
	}

	row, err := br.db.FindByID(context.Background(), br.col, ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	bBl := bsonBookLoan{}

	if row.Next() {
		if err := row.Scan(&bBl); err != nil {
			return nil, err
		}
	}

	if err := row.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return toModelBookLoan(&bBl), nil
}

// ListBookLoan fetches from repository
func (br *bookLoanRepository) ListBookLoan(ctx context.Context, skip *int64, limit *int64) ([]bookloan.BookLoan, error) {
	rows, err := br.db.Find(context.Background(), br.col, bson.M{}, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bBls := []bookloan.BookLoan{}
	for rows.Next() {
		bBk := bsonBookLoan{}
		if err := rows.Scan(&bBk); err != nil {
			return nil, err
		}
		bBls = append(bBls, *toModelBookLoan(&bBk))

	}
	if rows.Err(); err != nil {
		return nil, err
	}

	return bBls, nil
}

func (br *bookLoanRepository) ListBookLoanByUserID(ctx context.Context, userID string, skip *int64, limit *int64) ([]bookloan.BookLoan, error) {
	query := bson.M{
		"user_id": userID,
	}

	rows, err := br.db.Find(context.Background(), br.col, query, skip, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bBls := []bookloan.BookLoan{}
	for rows.Next() {
		bBl := bsonBookLoan{}
		if err := rows.Scan(&bBl); err != nil {
			return nil, err
		}
		bBls = append(bBls, *toModelBookLoan(&bBl))

	}
	if rows.Err(); err != nil {
		return nil, err
	}

	return bBls, nil
}

func (br *bookLoanRepository) UpdateBookLoan(ctx context.Context, ID string, bl *bookloan.BookLoan) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil
	}

	f := bson.M{
		"_id": objID,
	}

	now := time.Now()

	u := bson.M{
		"$set": bson.M{
			"book_id":         bl.BookID,
			"user_id":         bl.UserID,
			"status":          string(bl.Status),
			"accepted_at":     bl.AcceptedAt,
			"rejected_at":     bl.RejectedAt,
			"rejection_cause": bl.RejectionCause,
			"created_at":      bl.CreatedAt,
			"updated_at":      &now,
		},
	}

	res, err := br.db.UpdateOne(context.Background(), br.col, f, u)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
