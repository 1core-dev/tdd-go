package collection

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{3, 7, 8}

		got := Sum(numbers)
		want := 18

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 7}, []int{7, 13})
	want := []int{10, 20}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
