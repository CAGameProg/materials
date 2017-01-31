package main

import "fmt"

func main() {
	fmt.Print("Enter the amount of a contribution: ")
	var donation float64
	fmt.Scanf("%d", &donation)

	if donation >= 1000 {
		fmt.Print("You are a benefactor!")
	} else if donation >= 500 {
		fmt.Print("You are a patron!")
	} else if donation >= 100 {
		fmt.Print("You are a supporter!")
	} else if donation >= 15 {
		fmt.Print("You are a friend!")
	} else {
		fmt.Print("You are a cheapskate!")
	}
}
