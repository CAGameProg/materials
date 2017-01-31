package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter an integer: ")
	var max int
	fmt.Scanf("%d", &max)

	num := 0
	for i := 1; i <= max; i++ {
		root := int(math.Sqrt(float64(i)))
		if root*root == i {
			num++
		}
	}

	fmt.Println("There are", num, "perfect squares between 1 and", max)
}
