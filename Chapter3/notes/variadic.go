package main

import "fmt"

func Print(values ...int) {
	for i := 0; i < len(values); i++ {
		fmt.Print(values[i], " ")
	}
	fmt.Println()
}

func main() {
	Print(1, 2, 3, 4)
}
