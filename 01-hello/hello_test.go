package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Homer")
	want := "Hello, Homer"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
