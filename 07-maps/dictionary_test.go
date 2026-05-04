package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "test value"}

	got := Search(dict, "test")
	want := "test value"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
