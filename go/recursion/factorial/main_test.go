package main

import (
	"testing"
)

func TestFactorial0(t *testing.T) {
	got := factorial(0)
	want := 1

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestFactorial1(t *testing.T) {
	got := factorial(1)
	want := 1

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestFactorial5(t *testing.T) {
	got := factorial(5)
	want := 120

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}
