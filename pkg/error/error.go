package errwrapper

import (
	"errors"
	"fmt"
)

const (
	BadRequest        int = 400
	Unauthorized      int = 401
	InternalApiError  int = 404
	NotFound          int = 404
	RepositoryError   int = 500
	ThirdPartiesError int = 500
)

var (
	ErrorBadRequest = errors.New(`error bad request`)
)

type ErrorWrapper struct {
	ErrorCode    int
	ErrorMessage string
	Err          error
}

func (w *ErrorWrapper) Error() string {
	return fmt.Sprintf("%d: %v", w.ErrorCode, w.Err)
}

func ErrWrap(err error, errMessage string, errorCode int) *ErrorWrapper {
	return &ErrorWrapper{
		ErrorCode:    errorCode,
		ErrorMessage: errMessage,
		Err:          err,
	}
}
