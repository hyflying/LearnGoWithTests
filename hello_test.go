package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Evan")
	want := "Hello, Evan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
