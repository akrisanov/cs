package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1 // base case
	}
	return n * factorial(n-1)
}

func main() {
	number := 5
	fmt.Printf("Factorial of %d is %d\n", number, factorial(number))
}
