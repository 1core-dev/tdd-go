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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{8.0, 10.0}
		checkArea(t, rectangle, 80.0)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

}
