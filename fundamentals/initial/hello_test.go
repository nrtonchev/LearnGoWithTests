package main

import "testing"

func TestHello(t *testing.T) {
	get := Hello("Nik")
	want := "Hello, Nik"

	if get != want {
		t.Errorf("got %q want %q", get, want)
	}
}
