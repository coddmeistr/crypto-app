package app

import (
	"errors"
	"fmt"
	"net/http"
)

// Application errors.
var (
	// General
	ErrUnknown               = errors.New("Untracked error occured")
	ErrNotFound              = errors.New("Not found")
	ErrInternal              = errors.New("Internal")
	ErrBadRequest            = errors.New("Bad request")
	ErrValidation            = errors.New("Invalid request body. Check the documentation to know about all required fields and their types")
	ErrNotAllRequiredQueries = errors.New("Not all queries")

	// Ws
	ErrUnknownRequest = errors.New("Unknown request type")
	ErrUnknownEvent   = errors.New("Unknown event type")
	ErrUnknownChannel = errors.New("Unknown channel")

	// Export errors

)

var errorCodesMap = map[error]int{
	ErrNotFound:              404,
	ErrInternal:              500,
	ErrBadRequest:            400,
	ErrValidation:            3,
	ErrNotAllRequiredQueries: 5,
}

func WrapE(err error, msg string) error {
	return fmt.Errorf("%w. Info: "+msg, err)
}

func GetHTTPCodeFromError(err error) int {
	switch err {
	case ErrInternal:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrBadRequest, ErrValidation:
		return http.StatusBadRequest
	default:
		return http.StatusBadRequest
	}
}

func ErrorCode(err error) int {
	for k, v := range errorCodesMap {
		if errors.Is(err, k) {
			return v
		}
	}
	return 1
}
