package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep(secondsDuration time.Duration)
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep(secondsDuration time.Duration) {
	time.Sleep(secondsDuration)
}

func Countdown(out io.Writer, sleeper Sleeper, count int, finalWord string) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep(1 * time.Second)
	}
	fmt.Fprintln(out, finalWord)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper, 3, "Go!")
}
