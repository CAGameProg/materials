package main

import "fmt"

func average(nums []float64) float64 {
	total := 0.0

	for _, n := range nums {
		total += n
	}

	average := total / float64(len(nums))
	return average
}

func main() {
	nums := []float64{11.2, 86.5, 20, 13.2}
	fmt.Println(average(nums))
}
