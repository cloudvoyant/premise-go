package greeting_test

import (
	"fmt"

	greeting "example.com/premise-module/libs/premise-go-lib"
)

// Public API -----------------------------------------------------------------
func ExampleHello() {
	fmt.Println(greeting.Hello("Ada"))
	// Output: Hello, Ada!
}
