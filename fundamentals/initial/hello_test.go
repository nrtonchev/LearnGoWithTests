package main

import "testing"

func TestHello(t *testing.T) {
	get := Hello()
	want := "Hello, world"

	if get != want {
		t.Errorf("got %q want %q", get, want)
	}
}
