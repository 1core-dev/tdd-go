package structs

import "math"

type Rectangle struct {
	weight float64
	height float64
}

type Circle struct {
	radius float64
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
