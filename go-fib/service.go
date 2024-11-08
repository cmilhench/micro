package fib

import (
	_ "embed" // The "embed" package must be imported when using go:embed
	"fmt"
)

var Name string = "Fibonacci Service"

// The version number is calculated externally using go generate and then
// embedded within the binary using go embed
//
//go:generate sh -c "./scripts/version.sh > .version"
//go:embed .version
var Revision string

func ServiceName() string {
	return fmt.Sprintf("%s %s", Name, Revision)
}

// Below is the actual business logic "service". It's written like this to
// easily divorced it from network delivery medium such as HTTP, gRPC or, AMQP.
// For more information, look at hexagonal architecture, Although this is
// scratching the surface Of this design methodology

// Fibonacci returns the nth fibonacci number in the sequence
// conforming to the rule F(n) = F(n-1) + F(n-2)
// i.e. 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)

}

// Sequence returns a function that when called returns the next fibonacci
// number in the sequence conforming to the rule F(n) = F(n-1) + F(n-2)
// i.e. 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
func Sequence() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
