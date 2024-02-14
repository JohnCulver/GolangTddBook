package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep(duration time.Duration) {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}

func TestCountdown(t *testing.T) {
	const count = 3
	const finalWord = "Go!"

	t.Run("test string output from countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper, count, finalWord)
		got := buffer.String()
		want := `3
2
1
Go!
`
		if spySleeper.Calls != count {
			t.Errorf("Wanted %q calls, but got %q", count, spySleeper.Calls)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test proper operation order", func(t *testing.T) {
		sleeper := SpyCountdownOperations{}

		Countdown(&sleeper, &sleeper, count, finalWord)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleeper.Calls)
		}
	})

	t.Run("test sleep duration", func(t *testing.T) {
		timeSpy := SpyTime{}
		Countdown(os.Stdout, &timeSpy, count, finalWord)

		got := timeSpy.durationSlept.Seconds()
		want := float64(count)

		if got != want {
			t.Errorf("Slept for %f, but should have slept for %f", got, want)
		}
	})
}
