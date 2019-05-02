package errors

import "testing"

func TestNewHttpError(t *testing.T) {
	c := "some_code"
	m := "some_message"
	h := NewHttpError(400, NewError(c, m))
	if h.Status() != 400 || h.Code() != c || h.Message() != m {
		t.Error("Expects code, message, status are correct")
	}
}
