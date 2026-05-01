package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("got %q expected %q", repeated, expected)
	}
}
