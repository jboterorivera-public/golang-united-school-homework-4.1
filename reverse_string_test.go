package reverse_string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "Hello world!"
	want := "!dlrow olleH"
	got := ReverseString(s)

	if got != want {
		t.Errorf("TestReverseString = %v, want %v", got, want)
	}
}
