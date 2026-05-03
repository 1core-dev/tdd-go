package collection

import "testing"

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
