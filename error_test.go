package errors

import "testing"

func TestNewError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	e := NewError(c, m)
	if e.Code() != c || e.Message() != m {
		t.Error("Expect code, message are correct")
	}
}