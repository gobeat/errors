package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

type HttpError interface {
	error
	Code() string
	WithCode(code string) HttpError
	Message() string
	WithMessage(message string) HttpError
	Status() int
	WithStatus(status int) HttpError
	Cause() error
	WithCause(cause error) HttpError
}

// NewHttpError returns a HttpError
func NewHttpError(code string, message string, status int) HttpError {
	return &factoryHTTPError{
		code:    code,
		message: message,
		status:  status,
	}
}

type factoryHTTPError struct {
	code    string
	message string
	status  int
	cause   error
}

func (e *factoryHTTPError) Error() string {
	if e.cause != nil {
		return strings.TrimSpace(fmt.Sprintf("[%s] %s. Cause: %s", e.code, e.message, e.cause.Error()))
	}
	return strings.TrimSpace(fmt.Sprintf("[%s] %s", e.code, e.message))
}

func (e *factoryHTTPError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code,
		Message: e.message,
	})
}

func (e *factoryHTTPError) Code() string {
	return e.code
}

func (e *factoryHTTPError) WithCode(code string) HttpError {
	e.code = code
	return e
}

func (e *factoryHTTPError) Message() string {
	return e.message
}

func (e *factoryHTTPError) WithMessage(message string) HttpError {
	e.message = message
	return e
}

func (e *factoryHTTPError) Status() int {
	return e.status
}

func (e *factoryHTTPError) WithStatus(status int) HttpError {
	e.status = status
	return e
}

func (e *factoryHTTPError) Cause() error {
	return e.cause
}

func (e *factoryHTTPError) WithCause(cause error) HttpError {
	e.cause = cause
	return e
}
