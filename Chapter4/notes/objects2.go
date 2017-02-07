package main

import "fmt"

type Shape interface {
	Area() float64
	Circumference() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Circumference() float64 {
	return 3.14 * 2 * c.radius
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Circumference() float64 {
	return r.length*2 + r.width*2
}

func AreaPlusCircumference(s Shape) float64 {
	return s.Area() + s.Circumference()
}

func main() {
	rect := &Rectangle{3, 5}

	x := AreaPlusCircumference(rect)
	fmt.Println(x)
}
