package structure_interface

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64{
	return (t.Base * t.Height) * 0.5
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rect Rectangle) float64{
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64{
	return rect.Width * rect.Height
}


