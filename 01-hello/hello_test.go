package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Homer")
		want := "Hello, Homer"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sating 'Hello World' when no args provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
