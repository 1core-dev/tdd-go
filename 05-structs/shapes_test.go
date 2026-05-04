package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{15.0, 25.0}
		got := rectangle.Perimeter()
		want := 80.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		area Shape
		want float64
	}{
		{Rectangle{8.0, 10.0}, 80.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.area.Area()
		want := tt.want

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
}
