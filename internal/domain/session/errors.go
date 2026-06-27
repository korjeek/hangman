package session

import (
	"fmt"

	"github.com/google/uuid"
)

// ErrValidation represents an error that occurs when session initialization parameters
// fail to pass the business validation rules.
type ErrValidation struct {
	sessionId uuid.UUID
	msg       string
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf("Validation failed for session %s: %s", e.sessionId.String(), e.msg)
}
