package main

import (
	"testing"
)

func TestTreb(t *testing.T) {
	got, err := Part1("test.txt")
	if err != nil {
		t.Errorf("Part1(test.txt) returned error: %v", err)
	}

	want := 142

	if got != want {
		t.Errorf("Part1(test.txt) = %v, want %v", got, want)
	}
}

func TestTreb2(t *testing.T) {
	got, err := Part2("test2.txt")
	if err != nil {
		t.Errorf("Part2(test.txt) returned error: %v", err)
	}

	want := 281

	if got != want {
		t.Errorf("Part2(test.txt) = %v, want %v", got, want)
	}
}
