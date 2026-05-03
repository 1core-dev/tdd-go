package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{15.0, 25.0}
	got := Perimeter(rectangle)
	want := 80.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{8.0, 10.0}
	got := Area(rectangle)
	want := 80.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
