package errors

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	e := NewError(c, m)
	if e.Code() != c || e.Message() != m {
		t.Error("Expect code, message are correct")
	}
}

func TestErrorError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	e := NewError(c, m)
	expected := fmt.Sprintf("%s %s", c, m)
	actual := e.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Error("Unexpected error string")
	}

	_ = e.WithCode("")
	expected = fmt.Sprintf("%s", m)
	actual = e.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Error("Unexpected error string")
	}
}

func TestErrorJsonEncode(t *testing.T) {
	c := "some_code"
	m := "some_message"
	e := NewError(c, m)
	b, err := json.Marshal(e)
	j := "{\"code\":\"some_code\",\"message\":\"some_message\"}"
	if err != nil || strings.Compare(j, fmt.Sprintf("%s", b)) != 0 {
		t.Error("Unexpected json result")
	}
}
