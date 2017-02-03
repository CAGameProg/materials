package main

import "fmt"

func fibIter(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1

	for i := 0; i < n-1; i++ {
		tmp := b
		b = a + b
		a = tmp
	}

	return b
}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fibGenerator() func() int {
	a, b := 0, 1
	return func() int {
		n := a
		a, b = b, a+b
		return n
	}
}

func main() {
	fmt.Println(fib(40))
	fmt.Println(fibIter(40))

	nextFib := fibGenerator()
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
	fmt.Println(nextFib())
}
