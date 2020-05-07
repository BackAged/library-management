package rest

import (
	"encoding/json"
	"net/http"
)

// Response reponse serializer util
type Response struct {
	Status  int         `json:"-"`
	Message string      `json:"title,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// ServeJSON serves json to http client
func ServeJSON(status int, message string, data interface{}, errors interface{}, w http.ResponseWriter) error {
	resp := &Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}
