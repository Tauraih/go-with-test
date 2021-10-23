package main

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tau")

	got := buffer.String()
	want := "Hello, Tau"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}