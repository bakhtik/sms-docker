package main

import "testing"

func TestFoo(t *testing.T) {
	got := Foo()
	if got != "hola!" {
		t.Errorf("Got unexpected result: %v", got)
	}
}
