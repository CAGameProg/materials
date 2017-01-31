package main

import "fmt"

func main() {
	array := make([]int, 10)

	for i := 0; i < len(array); i++ {
		array[i] = i
	}

	array = append(array, 15)

	fmt.Println(array[2:])

	fmt.Println(array, len(array))
}
