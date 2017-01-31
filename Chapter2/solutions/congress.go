package main

import "fmt"

func main() {
	fmt.Print("Enter age of candidate: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Print("Enter years of U.S. citizenship: ")
	var citizenYears int
	fmt.Scanf("%d", &citizenYears)

	if age >= 30 && citizenYears >= 9 {
		fmt.Println("The candidate is eligible for election to both the House of Representatives and the Senate.")
	} else if age >= 25 && citizenYears >= 7 {
		fmt.Println("The candidate is eligible for election to the House of Representatives but is NOT eligible for election to the Senate.")
	} else {
		fmt.Println("The candidate is NOT eligible for election to either the House of Representatives or the Senate.")
	}
}
