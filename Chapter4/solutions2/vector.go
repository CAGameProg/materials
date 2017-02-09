package main

import (
	"fmt"
	"math"
)

func toDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}

type Vector struct {
	x, y float64
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vector) Angle() float64 {
	return toDegree(math.Atan2(v.y, v.x))
}

func (v *Vector) Mul(scalar float64) *Vector {
	return &Vector{v.x * scalar, v.y * scalar}
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{v.x + v2.x, v.y + v2.y}
}

func (v *Vector) Dot(v2 *Vector) float64 {
	return v.x*v2.x + v.y*v2.y
}

func (v *Vector) AngleBetween(v2 *Vector) float64 {
	dot := v.Dot(v2)
	return toDegree(math.Acos(dot / (v.Magnitude() * v2.Magnitude())))
}

func main() {
	v1 := &Vector{10, 2}
	v2 := &Vector{0, 5}

	fmt.Println(v1.Magnitude())
	fmt.Println(v1.Dot(v2))
	fmt.Println(v1.Mul(2))
	fmt.Println(v1.AngleBetween(v2))
}
