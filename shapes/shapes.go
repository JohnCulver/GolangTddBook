package shapes

import "math"

type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (tri Triangle) Area() float64 {
	return (tri.Base * tri.Height) * 0.5
}
