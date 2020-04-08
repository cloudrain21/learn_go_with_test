package structs

import "math"

func GetPerimeter(a, b float64) float64 {
	return 2 * a + 2 * b
}

func Area(w, h float64) float64 {
	return w * h
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Bottom float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle)Area() float64 {
	return r.Width * r.Height
}

func (c Circle)Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle)Area() float64 {
	return t.Bottom * t.Height / 2
}