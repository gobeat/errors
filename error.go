package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error interface {
	error
	WithError(err error) Error
	Code() string
	WithCode(code string) Error
	Message() string
	WithMessage(message string) Error
}

func NewError(code string, message string) Error {
	return &factoryError{
		code:    code,
		message: message,
	}
}

type factoryError struct {
	code    string
	message string
}

func (e *factoryError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code,
		Message: e.message,
	})
}

func (e *factoryError) WithError(err error) Error {
	return e.WithMessage(err.Error())
}

func (e *factoryError) Error() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", e.code, e.message))
}

func (e *factoryError) Code() string {
	return e.code
}

func (e *factoryError) WithCode(code string) Error {
	e.code = code
	return e
}

func (e *factoryError) Message() string {
	return e.message
}

func (e *factoryError) WithMessage(message string) Error {
	e.message = message
	return e
}
