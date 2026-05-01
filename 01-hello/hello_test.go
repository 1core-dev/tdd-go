package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Homer", "english")
		want := "Hello, Homer"

		assertCorrectMessage(t, got, want)
	})

	t.Run("sating 'Hello World' when no args provided", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Oliver", "spanish")
		want := "Hola, Oliver"

		assertCorrectMessage(t, got, want)
	})

}
