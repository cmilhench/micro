package fib

// Fibonacci returns the nth fibonacci number in the sequence
// conforming to the rule F(n) = F(n-1) + F(n-2)
// i.e. 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
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
