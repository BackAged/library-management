package rest

import (
	"net/http"

	"github.com/BackAged/library-management/book/domain/author"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/govalidator"
)

// NewAuthorRouter contains all routes for author service.
func NewAuthorRouter(h AuthorHandler) http.Handler {
	router := chi.NewRouter()
	router.With(AdminOnly).Post("/create", h.CreateAuthor)
	router.Get("/", h.ListAuthor)
	router.Get("/{authorID}", h.GetAuthor)

	return router
}

// AuthorHandler interface for the Book handlers.
type AuthorHandler interface {
	CreateAuthor(http.ResponseWriter, *http.Request)
	GetAuthor(http.ResponseWriter, *http.Request)
	ListAuthor(http.ResponseWriter, *http.Request)
}

type athrHandler struct {
	bkSvc author.Service
}

// NewAuthorHandler returns a new author handler
func NewAuthorHandler(bkSvc author.Service) AuthorHandler {
	return &athrHandler{bkSvc: bkSvc}
}

type createAuthorRequest struct {
	AuthorName string `json:"author_name"`
	Details    string `json:"details"`
}

type createAuthorReponse struct {
	ID         string `json:"id"`
	AuthorName string `json:"author_name"`
	Details    string `json:"details"`
}

func createAuthorRequestValidator(r *http.Request) (*createAuthorRequest, error) {
	var crtRq createAuthorRequest
	rules := govalidator.MapData{
		"author_name": []string{"required", "alpha_space"},
		"details":     []string{"required", "alpha_space"},
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
func (h *athrHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	crtRq, err := createAuthorRequestValidator(r)
	if err != nil {
		ServeJSON(http.StatusBadRequest, "", nil, err, w)
		return
	}

	bk := &author.Author{
		AuthorName: crtRq.AuthorName,
		Details:    crtRq.Details,
	}

	if err := h.bkSvc.Create(r.Context(), bk); err != nil {
		//  TODO:-> Domain level error handling
		ServeJSON(http.StatusInternalServerError, "error", nil, nil, w)
		return
	}

	resBk := &author.Author{
		ID:         bk.ID,
		AuthorName: bk.AuthorName,
		Details:    bk.Details,
	}
	ServeJSON(http.StatusCreated, "", resBk, nil, w)
}

type getAuthorReponse struct {
	ID         string `json:"id"`
	AuthorName string `json:"author_name"`
	Details    string `json:"details"`
}

// Get handler
func (h *athrHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "authorID")

	bk, err := h.bkSvc.Get(r.Context(), id)
	if err != nil {
		switch v := err.(type) {
		case *author.NotFound:
			ServeJSON(http.StatusBadRequest, v.GetMessage(), nil, v.GetErrors(), w)
		default:
			ServeJSON(http.StatusInternalServerError, "Something went wrong", nil, nil, w)
		}
		return
	}

	resBk := &getAuthorReponse{
		ID:         bk.ID,
		AuthorName: bk.AuthorName,
		Details:    bk.Details,
	}
	ServeJSON(http.StatusCreated, "", resBk, nil, w)
}

type listAuthorReponse struct {
	ID         string `json:"id"`
	AuthorName string `json:"author_name"`
	Details    string `json:"details"`
}

// List handler
func (h *athrHandler) ListAuthor(w http.ResponseWriter, r *http.Request) {
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

	resBks := []listAuthorReponse{}
	for _, b := range bk {
		resBk := listAuthorReponse{
			ID:         b.ID,
			AuthorName: b.AuthorName,
			Details:    b.Details,
		}
		resBks = append(resBks, resBk)
	}

	ServeJSON(http.StatusCreated, "", resBks, nil, w)
}
