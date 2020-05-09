package rest

import (
	"net/http"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/go-chi/chi"
	"github.com/thedevsaddam/govalidator"
)

// NewBookRouter contains all routes for books service.
func NewBookRouter(h BookHandler) http.Handler {
	router := chi.NewRouter()
	router.With(AdminOnly).Post("/create", h.CreateBook)
	router.Get("/", h.ListBook)
	router.Get("/{bookID}", h.GetBook)
	router.Get("/author/{authorID}", h.ListBookByAuthor)

	return router
}

// BookHandler interface for the Book handlers.
type BookHandler interface {
	CreateBook(http.ResponseWriter, *http.Request)
	GetBook(http.ResponseWriter, *http.Request)
	ListBook(http.ResponseWriter, *http.Request)
	ListBookByAuthor(http.ResponseWriter, *http.Request)
}

type bkHandler struct {
	bkSvc book.Service
}

// NewBookHandler will instantiate the handler
func NewBookHandler(bkSvc book.Service) BookHandler {
	return &bkHandler{bkSvc: bkSvc}
}

type createBookRequest struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	Quantity    int    `json:"quantity"`
}

type createBookReponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	Quantity    int    `json:"quantity"`
}

func createBookRequestValidator(r *http.Request) (*createBookRequest, error) {
	var crtRq createBookRequest
	rules := govalidator.MapData{
		"title":       []string{"required", "alpha_space"},
		"isbn":        []string{"required", "alpha_space"},
		"category":    []string{"required", "alpha_space"},
		"description": []string{"required", "alpha_space"},
		"author_id":   []string{"required", "alpha_space"},
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
func (h *bkHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	crtRq, err := createBookRequestValidator(r)
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
		Quantity:    crtRq.Quantity,
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

type getBookReponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Quantity    int    `json:"quantity"`
}

// Get handler
func (h *bkHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookID")

	bk, err := h.bkSvc.Get(r.Context(), id)
	if err != nil {
		switch v := err.(type) {
		case *book.NotFound:
			ServeJSON(http.StatusBadRequest, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	resBk := &getBookReponse{
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

type listBookReponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Quantity    int    `json:"quantity"`
}

// List handler
func (h *bkHandler) ListBook(w http.ResponseWriter, r *http.Request) {
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

	resBks := []listBookReponse{}
	for _, b := range bk {
		resBk := listBookReponse{
			ID:          b.ID,
			Title:       b.Title,
			ISBN:        b.ISBN,
			Category:    b.Category,
			Description: b.Description,
			AuthorID:    b.AuthorID,
			AuthorName:  b.AuthorName,
			Quantity:    b.Quantity,
		}
		resBks = append(resBks, resBk)
	}

	ServeJSON(http.StatusCreated, "", resBks, nil, w)
}

type listBookByAuthorReponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	Quantity    int    `json:"quantity"`
}

// List handler
func (h *bkHandler) ListBookByAuthor(w http.ResponseWriter, r *http.Request) {
	v, err := ParseSkipLimit(r)
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	skip, limit := v["skip"], v["limit"]
	authorID := chi.URLParam(r, "authorID")

	bk, err := h.bkSvc.GetAuthorBooks(r.Context(), authorID, &skip, &limit)
	if err != nil {
		//  TODO:-> Domain level error handling
		ServeJSON(http.StatusInternalServerError, "error", nil, nil, w)
		return
	}

	resBks := []listBookByAuthorReponse{}
	for _, b := range bk {
		resBk := listBookByAuthorReponse{
			ID:          b.ID,
			Title:       b.Title,
			ISBN:        b.ISBN,
			Category:    b.Category,
			Description: b.Description,
			AuthorID:    b.AuthorID,
			AuthorName:  b.AuthorName,
			Quantity:    b.Quantity,
		}
		resBks = append(resBks, resBk)
	}

	ServeJSON(http.StatusCreated, "", resBks, nil, w)
}
