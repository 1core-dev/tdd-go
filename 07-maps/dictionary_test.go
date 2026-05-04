package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got := dict.Search("test")
		want := "test value"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got := dict.Search("unknown")
		want := ""

		assertString(t, got, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
