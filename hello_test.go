package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		assertCorrectMessage(got, want, t)
	})
	t.Run("Say 'Hello, World' when an empty string is applied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
