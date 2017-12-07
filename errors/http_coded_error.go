package errors

import (
	"fmt"
	"net/http"
)

type CodedError interface {
	error
	Code() int
	Desc() string
}

type httpCodedError struct {
	code int
	desc string
}

var _ = httpCodedError{}

func NewHTTPCodedError(code int, format string, a ...interface{}) httpCodedError {
	return httpCodedError{code: code, desc: fmt.Sprintf(format, a)}
}

func NewHTTPNotFoundError(format string, a ...interface{}) httpCodedError {
	return NewHTTPCodedError(http.StatusNotFound, format, a)
}

func NewHTTPBadRequestError(format string, a ...interface{}) httpCodedError {
	return NewHTTPCodedError(http.StatusBadRequest, format, a)
}

func (e httpCodedError) Code() int {
	return e.code
}

func (e httpCodedError) Desc() string {
	return e.desc
}

func (e httpCodedError) Error() string {
	return fmt.Sprintf("status:%d, desc:%s", e.code, e.desc)
}
