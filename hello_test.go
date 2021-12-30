package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Alex")
	want := "Hello, Alex"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
