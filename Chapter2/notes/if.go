package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)

	if input == 10 {
		fmt.Println("The input was 10")
	} else if input == 11 {
		fmt.Println("The input was 11")
	} else {
		fmt.Println("The input was neither 10 nor 11")
	}

	fmt.Println("Done")
}
