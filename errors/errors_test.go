package errors

import (
	"encoding/json"
	"testing"
)

func Test_apierror_is_child_err(t *testing.T) {
	var err *APIError

	err = NewErrMissingParam("id")
	if !IsChildErr(err, ErrMissingParam) {
		t.Fail()
	}

	err = NewErrIncorrectParam("id")
	if !IsChildErr(err, ErrIncorrectParam) {
		t.Fail()
	}

	err = NewErrSQLNoRows("no rows")
	if !IsChildErr(err, ErrSQLNoRows) {
		t.Fail()
	}

	err = NewErrSQLInternalError("internal")
	if !IsChildErr(err, ErrSQLInternalError) {
		t.Fail()
	}
}

func Test_apierror_json(t *testing.T) {
	error := NewErrMissingParam("id")

	bs, err := json.Marshal(error)
	if err != nil {
		t.Fail()
	}

	error = &APIError{}
	if err := json.Unmarshal(bs, &error); err != nil {
		t.Fail()
	}

	if !IsChildErr(error, ErrMissingParam) {
		t.Fail()
	}
}
