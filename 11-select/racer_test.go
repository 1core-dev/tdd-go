package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://slowurl.com"
	fastURL := "https://fasturl.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
