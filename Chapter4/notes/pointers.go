package main

import "fmt"

func f(n *int) {
	*n = 5
	fmt.Println(*n)
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a, b := 5, 10
	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)
}
