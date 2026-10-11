package main

import (
	"bytes"
	"strings"
	"testing"
)

// Public API -----------------------------------------------------------------
func TestHelpAndVersion(t *testing.T) {
	for _, arg := range []string{"--help", "--version"} {
		c := command()
		var b bytes.Buffer
		c.SetOut(&b)
		c.SetArgs([]string{arg})
		if err := c.Execute(); err != nil {
			t.Fatal(err)
		}
		if strings.TrimSpace(b.String()) == "" {
			t.Fatal("empty output")
		}
	}
}
