package app

import (
	"errors"
	"fmt"
	"net/http"
)

// Application errors.
var (
	// General
	ErrNotFound              = errors.New("Not found")
	ErrInternal              = errors.New("Internal")
	ErrBadRequest            = errors.New("Bad request")
	ErrValidation            = errors.New("Invalid request body. Check the documentation to know about all required fields and their types")
	ErrInvalidParamType      = errors.New("Invalid param(s) type")
	ErrNotAllRequiredQueries = errors.New("Not all queries")

	// Export errors
	ErrIncorrectLoginOrPassword = errors.New("Incorrect login or password")
	ErrNotEnoughBalance         = errors.New("User's account has not enough balance to perform this operation")
)

var errorCodesMap = map[error]int{
	ErrNotFound:              404,
	ErrInternal:              500,
	ErrBadRequest:            400,
	ErrValidation:            3,
	ErrNotAllRequiredQueries: 5,

	ErrIncorrectLoginOrPassword: 11,
	ErrNotEnoughBalance:         12,
}

var codesToErrorsMap = map[int]error{
	404: ErrNotFound,
	500: ErrInternal,
	400: ErrBadRequest,
	3:   ErrValidation,
	5:   ErrNotAllRequiredQueries,

	11: ErrIncorrectLoginOrPassword,
	12: ErrNotEnoughBalance,
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
	return 1
}

func CodeToError(code int) (error, bool) {
	err, ok := codesToErrorsMap[code]
	return err, ok
}
