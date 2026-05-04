package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "test value"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
