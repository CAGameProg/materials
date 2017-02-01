package main

import "fmt"

func max(nums ...int) int {
	max := nums[0]

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func min(nums ...int) int {
	min := nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func main() {
	fmt.Println(max(5, 2, 10, 7))
	fmt.Println(min(5, 2, 10, 7))
}
