package main

import "fmt"

func Factorial(n uint) uint {
	if n <= 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func main() {
	fmt.Println(Factorial(100))
}
