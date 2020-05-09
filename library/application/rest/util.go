package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Response reponse serializer util
type Response struct {
	Status  int         `json:"-"`
	Message string      `json:"message,omitempty"`
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

// ParseSkipLimit parses query from query string
func ParseSkipLimit(r *http.Request) (map[string]int64, error) {
	m := make(map[string]int64)
	ve := &ValidationError{}

	s := r.URL.Query().Get("skip")
	if s != "" {
		skip, err := strconv.ParseInt(s, 10, 64)
		sErr := []string{}
		if err != nil {
			sErr = append(sErr, "skip must be a integer")
		} else {
			if skip < 0 {
				sErr = append(sErr, "skip must be a positive integer")
			}
		}

		if len(sErr) == 0 {
			m["skip"] = skip
		} else {
			ve.Add("skip", sErr)
		}
	} else {
		m["skip"] = 0
	}

	l := r.URL.Query().Get("limit")
	if l != "" {
		l := r.URL.Query().Get("limit")
		limit, err := strconv.ParseInt(l, 10, 64)
		lErr := []string{}
		if err != nil {
			lErr = append(lErr, "limit must be a integer")
		} else {
			if limit < 0 {
				lErr = append(lErr, "limit must be a positive integer")
			}

			if limit > 50 {
				limit = 25
			}
		}

		if len(lErr) == 0 {
			m["limit"] = limit
		} else {
			ve.Add("limit", lErr)
		}
	} else {
		m["limit"] = 50
	}

	if len(*ve) == 0 {
		return m, nil
	}

	return nil, ve
}
