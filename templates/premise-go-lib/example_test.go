package greeting_test

import (
	greeting "example.com/premise-module"
	"fmt"
)

// Public API -----------------------------------------------------------------
func ExampleHello() {
	fmt.Println(greeting.Hello("Ada"))
	// Output: Hello, Ada!
}
