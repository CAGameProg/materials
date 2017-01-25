package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var f float64
	fmt.Scanf("%f", &f)

	c := (f - 32) * 5 / 9

	fmt.Println(f, "degrees Fahrenheit is", c, "degrees Celsius")
}
