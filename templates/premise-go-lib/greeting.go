// Package greeting provides a small reusable greeting API.
package greeting

import "strings"

// Public API -----------------------------------------------------------------
// Hello returns a greeting, using world when name is empty.
func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "world"
	}
	return "Hello, " + name + "!"
}
