package structs

type Rectangle struct {
	weight float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.weight)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.weight
}
