package mocking

import (
	"bytes"
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(&buf, spySleeper)

		got := buf.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		sleepPrinter := &SpyCountdownOperations{}

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		Countdown(sleepPrinter, sleepPrinter)

		if !slices.Equal(want, sleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleepPrinter.Calls)
		}
	})
}

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write([]byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
