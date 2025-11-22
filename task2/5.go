package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l float64
	w float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.w
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.w)
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	r := Rectangle{2, 3}
	fmt.Printf("rectangle area = %v\n", r.Area())
	fmt.Printf("rectangle Perimeter = %v\n", r.Perimeter())

	c := Circle{2}
	fmt.Printf("Circle area = %v\n", c.Area())
	fmt.Printf("Circle Perimeter = %v\n", c.Perimeter())
}
