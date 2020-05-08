package author

import (
	"time"
)

// Author defines Author type
type Author struct {
	ID         string
	AuthorName string
	Details    string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (athr *Author) valid() (bool, error) {
	return true, nil
}
