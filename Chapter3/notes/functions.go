package main

import "fmt"

func Square(a int) int {
	return a * a
}

func main() {
	x := Square(5)
	fmt.Println(x)
}
