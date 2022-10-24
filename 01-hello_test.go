package main

import "testing"

func TestHello(t *testing.T) {
	expected := 1
	a := echo(expected)

	if a != expected {
		t.Errorf("error")
	}
}
