package main

import "testing"

// testing func Hello

// testing new branch

func TestHello(t *testing.T) {
	got := Hello("Mai")
	want := "Hello, Mai"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
