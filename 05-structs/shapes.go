package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	weight float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (r Rectangle) Area() float64 {
	return r.height * r.weight
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.weight)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.weight)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.weight
}
