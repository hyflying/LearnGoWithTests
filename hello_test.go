package main

import "testing"

func TestHello(t *testing.T) {
	// use subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run(("say 'Hello, world' when an empty string is supplied"), func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
