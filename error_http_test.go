package errors

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNewHttpError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	h := NewHttpError(c, m, 400)
	if h.Status() != 400 || h.Code() != c || h.Message() != m {
		t.Error("Expects code, message, status are correct")
	}
}

func TestHttpErrorJsonEncode(t *testing.T) {
	c := "some_code"
	m := "some_message"
	e := NewHttpError(c, m, 500)
	b, err := json.Marshal(e)
	j := "{\"code\":\"some_code\",\"message\":\"some_message\"}"
	if err != nil || strings.Compare(j, fmt.Sprintf("%s", b)) != 0 {
		t.Error("Unexpected json result")
	}
}

func TestNewHttpErrorImplementError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	he := NewHttpError(c, m, 500)
	_, ok := he.(error)
	if !ok {
		t.Error("HttpError must implement error")
	}
}
