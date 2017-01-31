package main

import "fmt"

func Square(n int) int {
	return n * n
}

func Double(n int) int {
	return n + n
}

func Compute(n int, computer func(int) int) int {
	return computer(n)
}

func main() {
	fmt.Println(Compute(5, Square))
	fmt.Println(Compute(5, Double))
}
