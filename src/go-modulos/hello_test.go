package hello

import (
	"testing"

	"rsc.io/quote"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
func Hello() string {
	return quote.Hello()
}
