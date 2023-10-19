package app

import (
	"errors"
	"fmt"
	"net/http"
)

type errorCode string

// Коды ошибок приложения.
var (
	// General
	ErrNotFound         = errors.New("Not found")
	ErrInternal         = errors.New("Internal")
	ErrBadRequest       = errors.New("Bad request")
	ErrValidation       = errors.New("Invalid request body. Check the documentation to know about all required fields and their types")
	ErrInvalidParamType = errors.New("Invalid param(s) type")

	// Export errors
	ErrIncorrectLoginOrPassword = errors.New("Incorrect login or password")
	ErrNotEnoughBalance         = errors.New("User's account has not enough balance to perform this operation")
)

var errorCodesMap = map[error]int{
	ErrNotFound:         404,
	ErrInternal:         405,
	ErrBadRequest:       406,
	ErrValidation:       407,
	ErrInvalidParamType: 408,

	ErrIncorrectLoginOrPassword: 500,
	ErrNotEnoughBalance:         501,
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
	case ErrBadRequest, ErrValidation, ErrInvalidParamType:
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
	return 400
}
