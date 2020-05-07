package rest

import "encoding/json"

// ValidationError Web validation error
type ValidationError map[string][]string

func (ve *ValidationError) Error() string {
	b, _ := json.Marshal(ve)
	return string(b)
}

// Add new
func (ve ValidationError) Add(key string, value []string) {
	ve[key] = value
}
