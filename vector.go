package main

import "math"

type Vector2 struct {
	X, Y float64
}

func (a Vector2) Add(b Vector2) Vector2 {
	return Vector2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector2) Sub(b Vector2) Vector2 {
	return Vector2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector2) MultiplyByScalar(s float64) Vector2 {
	return Vector2{
		X: a.X * s,
		Y: a.Y * s,
	}
}

func (a Vector2) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector2) Normalize() Vector2 {
	return a.MultiplyByScalar(1.0 / a.Length())
}

func (a Vector2) Dot(b Vector2) float64 {
	return (a.X * b.X) + (a.Y * b.Y)
}

func (a Vector2) Cross(b Vector2) float64 {
	return (a.X * b.Y) - (a.Y * b.X)
}
