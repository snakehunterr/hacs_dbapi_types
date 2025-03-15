package errors

import "fmt"

type APIErrorCode byte

const (
	ErrCodeMissingParam APIErrorCode = iota + 1
	ErrCodeIncorrectParam
	ErrCodeSQLNoRows
	ErrCodeSQLInternalError
)

type APIError struct {
	Code APIErrorCode `json:"code,omitempty"`
	Err  string       `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	if e == nil {
		return "*APIError is nil (no error)"
	}
	return e.Err
}

var (
	ErrMissingParam     = NewErrMissingParam("")
	ErrIncorrectParam   = NewErrIncorrectParam("")
	ErrSQLNoRows        = NewErrSQLNoRows("")
	ErrSQLInternalError = NewErrSQLInternalError("")
)

func NewErrMissingParam(param string) *APIError {
	return &APIError{
		Code: ErrCodeMissingParam,
		Err:  fmt.Sprintf("missing param: %s", param),
	}
}

func NewErrIncorrectParam(param string) *APIError {
	return &APIError{
		Code: ErrCodeIncorrectParam,
		Err:  fmt.Sprintf("incorrect param: %s", param),
	}
}

func NewErrSQLNoRows(err string) *APIError {
	return &APIError{
		Code: ErrCodeSQLNoRows,
		Err:  err,
	}
}

func NewErrSQLInternalError(err string) *APIError {
	return &APIError{
		Code: ErrCodeSQLInternalError,
		Err:  err,
	}
}

func IsChildErr(child, parent *APIError) bool {
	if child == nil || parent == nil {
		return false
	}

	return child.Code == parent.Code
}
