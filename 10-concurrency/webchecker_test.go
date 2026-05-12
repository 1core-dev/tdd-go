package concurrency

import (
	"maps"
	"testing"
)

func mockWebChecker(url string) bool {
	return url != "waat://invalid.site"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://gooole.com",
		"https://www.wikipedia.org",
		"waat://invalid.site",
	}

	want := map[string]bool{
		"https://gooole.com":        true,
		"https://www.wikipedia.org": true,
		"waat://invalid.site":       false,
	}

	got := CheckWebsites(mockWebChecker, urls)

	if !maps.Equal(got, want) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}
