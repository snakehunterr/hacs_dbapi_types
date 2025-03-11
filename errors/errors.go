package errors

type APIErrorCode byte

const (
	ErrMissingParam APIErrorCode = iota
	ErrIncorrectParam
	ErrSQLNoRows
	ErrSQLInternalError
)
