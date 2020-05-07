package rest

import (
	"fmt"
	"net/http"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/go-chi/chi"
	"github.com/thedevsaddam/govalidator"
)

// NewBookRouter contains all routes for books service.
func NewBookRouter(h BookHandler) http.Handler {
	router := chi.NewRouter()
	router.Post("/create", h.Create)

	return router
}

// BookHandler interface for the Book handlers.
type BookHandler interface {
	Create(http.ResponseWriter, *http.Request)
}

type bkHandler struct {
	bkSvc book.Service
}

// NewBookHandler will instantiate the handler
func NewBookHandler(bkSvc book.Service) BookHandler {
	return &bkHandler{bkSvc: bkSvc}
}

type createRequest struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Quantity    int    `json:"quantity"`
}

type createReponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Quantity    int    `json:"quantity"`
}

func createRequestValidator(r *http.Request) (*createRequest, error) {
	var crtRq createRequest
	rules := govalidator.MapData{
		"title":       []string{"required", "alpha_space"},
		"isbn":        []string{"required", "alpha_space"},
		"category":    []string{"required", "alpha_space"},
		"description": []string{"required", "alpha_space"},
		"author_id":   []string{"required", "alpha_space"},
		"author_name": []string{"required", "alpha_space"},
		"quantity":    []string{"required", "numeric"},
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
func (h *bkHandler) Create(w http.ResponseWriter, r *http.Request) {
	crtRq, err := createRequestValidator(r)
	fmt.Println(crtRq, err)
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	bk := &book.Book{
		Title:       crtRq.Title,
		ISBN:        crtRq.ISBN,
		Category:    crtRq.Category,
		Description: crtRq.Description,
		AuthorID:    crtRq.AuthorID,
		AuthorName:  crtRq.AuthorName,
		Quantity:    crtRq.Quantity,
	}

	if err := h.bkSvc.Create(r.Context(), bk); err != nil {
		//  TODO:-> Domain level error handling
		ServeJSON(http.StatusInternalServerError, "error", nil, nil, w)
		return
	}

	resBk := &book.Book{
		ID:          bk.ID,
		Title:       bk.Title,
		ISBN:        bk.ISBN,
		Category:    bk.Category,
		Description: bk.Description,
		AuthorID:    bk.AuthorID,
		AuthorName:  bk.AuthorName,
		Quantity:    bk.Quantity,
	}
	ServeJSON(http.StatusCreated, "", resBk, nil, w)
}
