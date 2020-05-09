package rest

import (
	"net/http"
	"time"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/BackAged/library-management/library/domain/bookloan"
	"github.com/go-chi/chi"
	"github.com/thedevsaddam/govalidator"
)

// NewBookLoanRouter contains all routes for books loan service.
func NewBookLoanRouter(h BookLoanHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/create", h.CreateBookLoan)
	router.Get("/", h.ListBookLoan)
	router.Get("/{bookLoanID}", h.GetBookLoan)
	router.With(AdminOnly).Get("/{bookLoanID}/accept", h.AcceptBookLoan)
	router.With(AdminOnly).Post("/{bookLoanID}/reject", h.RejectBookLoan)

	return router
}

// BookLoanHandler interface for the BookLoan handlers.
type BookLoanHandler interface {
	CreateBookLoan(http.ResponseWriter, *http.Request)
	GetBookLoan(http.ResponseWriter, *http.Request)
	ListBookLoan(http.ResponseWriter, *http.Request)
	AcceptBookLoan(http.ResponseWriter, *http.Request)
	RejectBookLoan(http.ResponseWriter, *http.Request)
}

type bklnHandler struct {
	bkSvc bookloan.Service
}

// NewBookLoanHandler will instantiate the handler
func NewBookLoanHandler(bkSvc bookloan.Service) BookLoanHandler {
	return &bklnHandler{bkSvc: bkSvc}
}

type createBookLoanRequest struct {
	BookID string `json:"book_id"`
	UserID string `json:"user_id"`
}

type createBookLoanReponse struct {
	ID     string `json:"id"`
	BookID string `json:"book_id"`
	UserID string `json:"user_id"`
	Status string `json:"status"`
}

func createBookLoanRequestValidator(r *http.Request) (*createBookLoanRequest, error) {
	var crtRq createBookLoanRequest
	rules := govalidator.MapData{
		"book_id": []string{"required", "alpha_space"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &crtRq,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	err := v.ValidateJSON()
	if len(err) == 0 {
		return &crtRq, nil
	}

	ve := &ValidationError{}
	for k, v := range err {
		ve.Add(k, v)
	}

	return nil, ve
}

// Create handler
func (h *bklnHandler) CreateBookLoan(w http.ResponseWriter, r *http.Request) {
	crtRq, err := createBookLoanRequestValidator(r)
	userID := r.Header.Get("x-userid")
	if userID == "" {
		ServeJSON(http.StatusForbidden, "Un authorized", nil, nil, w)
		return
	}
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	bk := &bookloan.BookLoan{
		BookID: crtRq.BookID,
		UserID: userID,
	}

	if err := h.bkSvc.Create(r.Context(), bk); err != nil {
		switch v := err.(type) {
		case *book.AuthorNotFound:
			ServeJSON(http.StatusInternalServerError, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	resBk := &createBookLoanReponse{
		ID:     bk.ID,
		BookID: bk.BookID,
		UserID: bk.UserID,
		Status: string(bk.Status),
	}
	ServeJSON(http.StatusCreated, "", resBk, nil, w)
}

type getBookLoanReponse struct {
	ID             string     `json:"id"`
	BookID         string     `json:"book_id"`
	UserID         string     `json:"user_id"`
	Status         string     `json:"status"`
	AcceptedAt     *time.Time `json:"accepted_at,omitempty"`
	RejectedAt     *time.Time `json:"rejected_at,omitempty"`
	RejectionCause string     `json:"rejection_cause,omitempty"`
}

// GetBookLoan handler
func (h *bklnHandler) GetBookLoan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookLoanID")

	bk, err := h.bkSvc.Get(r.Context(), id)
	if err != nil {
		switch v := err.(type) {
		case *bookloan.NotFound:
			ServeJSON(http.StatusBadRequest, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	resBk := &getBookLoanReponse{
		ID:             bk.ID,
		BookID:         bk.BookID,
		UserID:         bk.UserID,
		Status:         string(bk.Status),
		AcceptedAt:     bk.AcceptedAt,
		RejectedAt:     bk.RejectedAt,
		RejectionCause: bk.RejectionCause,
	}
	ServeJSON(http.StatusCreated, "", resBk, nil, w)
}

type listBookLoanReponse struct {
	ID             string     `json:"id"`
	BookID         string     `json:"book_id"`
	UserID         string     `json:"user_id"`
	Status         string     `json:"status"`
	AcceptedAt     *time.Time `json:"accepted_at,omitempty"`
	RejectedAt     *time.Time `json:"rejected_at,omitempty"`
	RejectionCause string     `json:"rejection_cause,omitempty"`
}

// List handler
func (h *bklnHandler) ListBookLoan(w http.ResponseWriter, r *http.Request) {
	v, err := ParseSkipLimit(r)
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	skip, limit := v["skip"], v["limit"]

	bk, err := h.bkSvc.List(r.Context(), &skip, &limit)
	if err != nil {
		//  TODO:-> Domain level error handling
		ServeJSON(http.StatusInternalServerError, "error", nil, nil, w)
		return
	}

	resBks := []listBookLoanReponse{}
	for _, b := range bk {
		resBk := listBookLoanReponse{
			ID:             b.ID,
			BookID:         b.BookID,
			UserID:         b.UserID,
			Status:         string(b.Status),
			AcceptedAt:     b.AcceptedAt,
			RejectedAt:     b.RejectedAt,
			RejectionCause: b.RejectionCause,
		}
		resBks = append(resBks, resBk)
	}

	ServeJSON(http.StatusCreated, "", resBks, nil, w)
}

// List handler
func (h *bklnHandler) AcceptBookLoan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookLoanID")

	err := h.bkSvc.Accept(r.Context(), id)
	if err != nil {
		switch v := err.(type) {
		case *bookloan.NotFound:
			ServeJSON(http.StatusBadRequest, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	ServeJSON(http.StatusCreated, "book loan was accepted", nil, nil, w)
}

type rejectBookLoanRequest struct {
	RejectionCause string `json:"rejection_cause"`
}

func rejectBookLoanRequestValidator(r *http.Request) (*rejectBookLoanRequest, error) {
	var crtRq rejectBookLoanRequest
	rules := govalidator.MapData{
		"rejection_cause": []string{"required", "alpha_space"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &crtRq,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	err := v.ValidateJSON()
	if len(err) == 0 {
		return &crtRq, nil
	}

	ve := &ValidationError{}
	for k, v := range err {
		ve.Add(k, v)
	}

	return nil, ve
}

// List handler
func (h *bklnHandler) RejectBookLoan(w http.ResponseWriter, r *http.Request) {
	crtRq, err := rejectBookLoanRequestValidator(r)
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	id := chi.URLParam(r, "bookLoanID")

	err = h.bkSvc.Reject(r.Context(), id, crtRq.RejectionCause)
	if err != nil {
		switch v := err.(type) {
		case *bookloan.NotFound:
			ServeJSON(http.StatusBadRequest, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	ServeJSON(http.StatusCreated, "book loan was rejected", nil, nil, w)
}
