package main

import (
	"fmt"
)

// Fib returns nth(from 0th) Fibonacci number
func Fib(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}

func main() {
	fmt.Println(Fib(5))
}
