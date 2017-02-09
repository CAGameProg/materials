package main

import "fmt"

type Color struct {
	r byte
	g byte
	b byte
}

type Vehicle struct {
	maxSpeed int
	brand    string
	color    Color
}

func (v *Vehicle) Repaint(newColor Color) {
	v.color = newColor
}

type Car struct {
	*Vehicle

	seats    int
	plateNum string
}

func main() {
	car := &Car{&Vehicle{60, "ford", Color{255, 0, 0}}, 4, "10X3RH"}
	fmt.Println(car.color)

	car.Repaint(Color{0, 255, 0})
	fmt.Println(car.color)
}
