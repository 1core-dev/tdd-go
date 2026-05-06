package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
