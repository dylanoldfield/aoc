package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 13
	got, err := Part1("test_input.txt")

	if err != nil {
		t.Errorf("Part 1 failed with err: %v", err)
	}

	if got != want {
		t.Errorf("scores don't match. wanted:%v, got:%v", want, got)
	}

}

func TestPart2(t *testing.T) {
	want := 30
	got, err := Part2("test_input.txt")

	if err != nil {
		t.Errorf("Part 2 failed with err: %v", err)
	}

	if got != want {
		t.Errorf("scores don't match. wanted:%v, got:%v", want, got)
	}
}
