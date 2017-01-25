package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter the size of the integer: ")
	var size int
	fmt.Scanf("%d", &size)

	maxNum := math.Pow(2.0, float64(size)) - 1

	fmt.Println("An integer with", size, "bits of storage has a max of", maxNum)
}
