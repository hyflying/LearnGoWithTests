package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
