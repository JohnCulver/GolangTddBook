package shapes

import "math"

type Shape interface {
	Area() float64
	Name() string
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Floor(2 * math.Pi * c.Radius)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (tri Triangle) Area() float64 {
	return (tri.Base * tri.Height) * 0.5
}

func (t Triangle) Perimeter() float64 {
	return -1
}

func (t Triangle) Name() string {
	return "Triangle"
}
