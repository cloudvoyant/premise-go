package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

// Public API -----------------------------------------------------------------
func TestInteraction(t *testing.T) {
	s := httptest.NewServer(handler())
	defer s.Close()
	r, err := s.Client().Post(s.URL+"/greet", "application/x-www-form-urlencoded", strings.NewReader("name=%3Cscript%3E"))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	if !strings.Contains(string(b), "&lt;script&gt;") {
		t.Fatal(string(b))
	}
}
