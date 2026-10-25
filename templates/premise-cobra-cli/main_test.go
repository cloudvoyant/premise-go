package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// Fixtures -------------------------------------------------------------------

var errOutput = errors.New("output unavailable")

type failingWriter struct{}

// Tests ----------------------------------------------------------------------

func TestCommand(t *testing.T) {
	for _, argument := range []string{"--help", "--version"} {
		t.Run("when requesting "+argument, func(t *testing.T) {
			// Arrange
			command := command()
			var output bytes.Buffer
			command.SetOut(&output)
			command.SetArgs([]string{argument})

			// Act
			err := command.Execute()

			// Assert
			t.Run("it writes the requested information", func(t *testing.T) {
				if err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(output.String()) == "" {
					t.Fatal("empty output")
				}
			})
		})
	}
	t.Run("when writing the greeting fails", func(t *testing.T) {
		// Arrange
		command := command()
		command.SetOut(failingWriter{})
		command.SetArgs([]string{})

		// Act
		err := command.Execute()

		// Assert
		t.Run("it reports the output failure", func(t *testing.T) {
			if !errors.Is(err, errOutput) {
				t.Fatalf("error = %v, want %v", err, errOutput)
			}
		})
	})
}

// Private helpers ------------------------------------------------------------

func (failingWriter) Write([]byte) (int, error) { return 0, errOutput }
