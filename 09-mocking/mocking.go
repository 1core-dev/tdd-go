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

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
