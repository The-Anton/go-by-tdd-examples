package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main () {
	fmt.Println(Rectangle{20.0, 40.0})
}