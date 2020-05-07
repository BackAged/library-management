package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BackAged/library-management/book/domain/book"
	"github.com/go-chi/chi"
)

// BookRouter contains all routes for albums service.
func BookRouter(h BookHandler) http.Handler {
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
func NewBookHandler(tskSvc book.Service) BookHandler {
	return &bkHandler{bkSvc: tskSvc}
}

type createDTO struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id"`
}

// Create handler
func (h *bkHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("shit")
	tskDTO := &createDTO{}
	if err := json.NewDecoder(r.Body).Decode(&tskDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tsk := &book.Book{
		Title:       tskDTO.Title,
		Category:    tskDTO.Category,
		Description: tskDTO.Description,
		AuthorID:    tskDTO.AuthorID,
	}

	fmt.Println(tsk)

	if err := h.bkSvc.Create(r.Context(), tsk); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(tsk)
	json.NewEncoder(w).Encode(tsk.ID)
}
