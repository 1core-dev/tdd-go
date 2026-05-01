package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in English", func(t *testing.T) {
		got := Hello("Homer", "english")
		want := "Hello, Homer"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in English with empty name", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Oliver", "Spanish")
		want := "Hola, Oliver"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Oliver", "French")
		want := "Bonjour, Oliver"

		assertCorrectMessage(t, got, want)
	})

}
