package main

import (
	"math"
)

type Shape interface {
	Type() string
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Type() string {
	return "Circle"
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Type() string {
	return "Rectangle"
}
