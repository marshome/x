package errors

import (
	"net/http"
	"encoding/json"
)

const (
	INTERNAL_EXCEPTION = "InternalException"
	INVALID_ARGS = "InvalidArgs"
)

type Error struct {
	HttpStatus int
	Message    string
	Errors     []*ErrorItem
}

type ErrorItem struct {
	Reason  string
	Message string
}

func (e *Error)Error()string {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func NewError(message string) *Error {
	return &Error{
		HttpStatus:http.StatusBadRequest,
		Message:message,
	}
}

func NewInternalException(message string) *Error {
	return &Error{
		HttpStatus:http.StatusInternalServerError,
		Message:message,
	}
}

func NewInvalidArgs(message string, errors...*ErrorItem) *Error {
	e := &Error{
		HttpStatus:http.StatusBadRequest,
		Message:message,
	}

	if errors != nil {
		e.Errors = append(e.Errors, errors...)
	}

	return e
}