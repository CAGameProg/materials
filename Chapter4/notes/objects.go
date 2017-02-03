package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Zero(p1 *Point) {
	p1.x = 0
	p1.y = 0
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	p1 := &Point{10, 5}
	p2 := &Point{7, 0}

	fmt.Println(Distance(p1, p2))

	Zero(p1)
	fmt.Println(p1)
}
