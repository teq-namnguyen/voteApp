package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type TeqError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
	IsSentry  bool
}

func (e TeqError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

func (e TeqError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string, isSentry bool) TeqError {
	return TeqError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
		IsSentry:  isSentry,
	}
}

func ErrExampleGet(err error) TeqError {
	return TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10000",
		Message:   "Failed to get example.",
		IsSentry:  true,
	}
}

func ErrGetWithString() TeqError {
	return TeqError{
		Raw:       nil,
		ErrorCode: "10000",
		HTTPCode:  http.StatusBadRequest,
		Message:   "Wrong password",
		IsSentry:  false,
	}
}

func ErrExampleCreate(err error) TeqError {
	return TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10001",
		Message:   "Failed to create example.",
		IsSentry:  true,
	}
}

func ErrExampleUpdate(err error) TeqError {
	return TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10002",
		Message:   "Failed to update example.",
		IsSentry:  true,
	}
}

func ErrExampleDelete(err error) TeqError {
	return TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10003",
		Message:   "Failed to delete example.",
		IsSentry:  true,
	}
}

func ErrExampleNotFound() TeqError {
	return TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrExampleInvalidParam(param string) TeqError {
	return TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrToken(messErr string) TeqError {
	return TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10006",
		Message:   messErr,
		IsSentry:  false,
	}
}

func ErrServerInternals(messErr string) TeqError {
	return TeqError{
		Raw:       nil,
		ErrorCode: "10007",
		HTTPCode:  500,
		Message:   messErr,
		IsSentry:  false,
	}
}
