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

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "test value"

		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test value"
		dict := Dictionary{word: definition}

		err := dict.Add(word, "new value")
		assertError(t, err, ErrWordExist)
		assertDefinition(t, dict, word, definition)
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

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, definition)
}
