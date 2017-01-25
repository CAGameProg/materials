package main

import "fmt"

func main() {
	var array [1000]int

	for i := 0; i < len(array); i++ {
		array[i] = i
	}

	fmt.Println(array)
	// Should print [0 1 2 3 4 5 6 7 8 ... 999]
}
