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

func (e APIError) Error() string {
	return e.Err
}

var (
	ErrMissingParam     = NewErrMissingParam("")
	ErrIncorrectParam   = NewErrIncorrectParam("")
	ErrSQLNoRows        = APIError{Code: ErrCodeSQLNoRows, Err: "no rows from sql query"}
	ErrSQLInternalError = APIError{Code: ErrCodeSQLInternalError, Err: "sql server internal error"}
)

func NewErrMissingParam(param string) APIError {
	return APIError{
		Code: ErrCodeMissingParam,
		Err:  fmt.Sprintf("missing param: %s", param),
	}
}

func NewErrIncorrectParam(param string) APIError {
	return APIError{
		Code: ErrCodeIncorrectParam,
		Err:  fmt.Sprintf("incorrect param: %s", param),
	}
}

func IsChildErr(child, parent APIError) bool {
	return child.Code == parent.Code
}
