package inwx

import (
	"fmt"
)

// ErrorCode represents an error code returned from the API.
type ErrorCode int

// Error codes returned from the API.
const (
	// AuthenticationFailedError gets thrown for wrong credentials.
	AuthenticationFailedError ErrorCode = 2200
)

// An Error reports the error caused by an API request.
type Error struct {
	Status  int
	Message string
	Code    string
	Reason  string
}

// Error implements the standard lib error function.
func (e Error) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf(
			"(%d) %s, reason: (%s) %s",
			e.Status,
			e.Message,
			e.Code,
			e.Reason,
		)
	}

	return fmt.Sprintf(
		"(%d) %s",
		e.Status,
		e.Message,
	)
}

// IsError returns whether err is an API error with the given error code.
func IsError(err error, status ErrorCode) bool {
	apiErr, ok := err.(Error)
	return ok && apiErr.Status == int(status)
}
