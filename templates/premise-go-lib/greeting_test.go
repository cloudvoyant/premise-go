package greeting

import "testing"

// Public API -----------------------------------------------------------------
func TestHello(t *testing.T) {
	for in, want := range map[string]string{"": "Hello, world!", " Ada ": "Hello, Ada!"} {
		if got := Hello(in); got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}
}
