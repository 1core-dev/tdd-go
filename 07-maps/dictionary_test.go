package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test value"}

	got := dict.Search("test")
	want := "test value"

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
