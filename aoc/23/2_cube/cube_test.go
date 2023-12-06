package main

import (
	"testing"
)

func TestGameNumber(t *testing.T) {
	base := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	wantNum := 1
	wantRemainder := "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	gotNum, gotRemainder, err := GetGameNumber(base)

	if err != nil {
		t.Errorf("unable to get game number: %v", err)
	}

	if gotNum != wantNum {
		t.Errorf("Numbers don't match. Got: %v, Wanted: %v", gotNum, wantNum)
	}

	if wantRemainder != gotRemainder {
		t.Errorf("Remainders don't match. Got: %v, Wanted: %v", gotRemainder, wantRemainder)
	}
}

func TestGetBlockPairs(t *testing.T) {
	base := "3 blue, 4 red, 2 green"
	var want [3]BlockPair

	want[0] = BlockPair{3, "blue"}
	want[1] = BlockPair{4, "red"}
	want[2] = BlockPair{2, "green"}

	gotPairs, err := GetBlockPairs(base)

	if err != nil {
		t.Errorf("failed to get pairs %v", err)
	}

	for i, pair := range gotPairs {
		if pair.Count != want[i].Count || pair.Colour != want[i].Colour {
			t.Errorf("Pair doesn't match, wanted: %v got: %v", want[i], pair)
		}
	}
}

func TestPart1(t *testing.T) {
	base := "test1.txt"
	want := 8
	got, err := Part1(base)

	if err != nil {
		t.Errorf("failure in part 1, %v", err)
	}

	if want != got {
		t.Errorf("failure in part 1, got: %v, want: %v", got, want)
	}
}
