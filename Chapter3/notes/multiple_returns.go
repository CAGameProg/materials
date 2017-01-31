package main

import "fmt"

func MultipleSquare(a, b int) (int, int) {
	return a * a, b * b
}

func main() {
	a, b := MultipleSquare(5, 4)

	fmt.Println(a)
	fmt.Println(b)
}
