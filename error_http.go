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
}

func NewHttpError(code string, message string, status int) HttpError {
	return &factoryHttpError{
		code:    code,
		message: message,
		status:  status,
	}
}

type factoryHttpError struct {
	code    string
	message string
	status  int
}

func (e *factoryHttpError) Error() string {
	return strings.TrimSpace(fmt.Sprintf("[%s] %s", e.code, e.message))
}

func (e *factoryHttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code,
		Message: e.message,
	})
}

func (e *factoryHttpError) Code() string {
	return e.code
}

func (e *factoryHttpError) WithCode(code string) HttpError {
	e.code = code
	return e
}

func (e *factoryHttpError) Message() string {
	return e.message
}

func (e *factoryHttpError) WithMessage(message string) HttpError {
	e.message = message
	return e
}

func (e *factoryHttpError) Status() int {
	return e.status
}

func (e *factoryHttpError) WithStatus(status int) HttpError {
	e.status = status
	return e
}
