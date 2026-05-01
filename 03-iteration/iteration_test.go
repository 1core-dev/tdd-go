package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("got %q expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 8)
	fmt.Println(repeated)
	// Output: bbbbbbbb
}
