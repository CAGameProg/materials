package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x      float64
	y      float64
	radius float64
}

func MakeUnitCircle() *Circle {
	return &Circle{0, 0, 1}
}

func CircleArea(c *Circle) float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	fmt.Println(CircleArea(MakeUnitCircle()))
}
