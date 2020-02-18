package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	circle := circle{0, 0, 5}
	rectangle := rectangle{10, 5}
	// fmt.Println(getArea(circle), getArea(rectangle))
	fmt.Println(circle.area(), rectangle.area())
}
