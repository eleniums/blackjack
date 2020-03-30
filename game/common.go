package game

import (
	"github.com/gofrs/uuid"
)

// NewUUID safely creates a UUID string.
func NewUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
