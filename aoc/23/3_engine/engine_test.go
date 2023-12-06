package main

import (
	"reflect"
	"testing"
)

func TestLoadSchematic(t *testing.T) {
	base := "schema_test.txt"
	var wantSMap = make(map[int][]int)
	wantSMap[1] = []int{2}

	var wantNMap = make(map[int]map[int]NumberPosition)
	wantNMap[0] = make(map[int]NumberPosition)
	wantNMap[0][1] = NumberPosition{23, Range{1, 2}}
	wantNMap[0][2] = NumberPosition{23, Range{1, 2}}

	wantNMap[2] = make(map[int]NumberPosition)
	wantNMap[2][0] = NumberPosition{1, Range{0, 0}}

	wantNMap[2][2] = NumberPosition{12, Range{2, 3}}
	wantNMap[2][3] = NumberPosition{12, Range{2, 3}}

	wantBoundary := Boundary{2, 4}
	gotSMap, gotNMap, gotBoundary, err := LoadSchematic(base, '.', false)

	if err != nil {
		t.Errorf("failed to load schema, %v", err)
	}

	if !reflect.DeepEqual(wantSMap, gotSMap) {
		t.Errorf("symbol maps dont match, got: %v, want: %v", gotSMap, wantSMap)
	}

	if !reflect.DeepEqual(wantNMap, gotNMap) {
		t.Errorf("number maps dont match, got: %v, want: %v", gotNMap, wantNMap)
	}

	if !reflect.DeepEqual(wantBoundary, gotBoundary) {
		t.Errorf("File Sizes is not what expected. got %v, want: %v", gotBoundary, wantBoundary)
	}
}

func TestCheckandRemoveNum(t *testing.T) {
	baseNMap := make(map[int]map[int]NumberPosition)

	baseNMap[0] = make(map[int]NumberPosition)
	baseNMap[0][1] = NumberPosition{23, Range{0, 2}}
	baseNMap[0][2] = NumberPosition{23, Range{0, 2}}

	baseNMap[2] = make(map[int]NumberPosition)
	baseNMap[2][0] = NumberPosition{1, Range{0, 0}}

	baseNMap[2][2] = NumberPosition{12, Range{2, 3}}
	baseNMap[2][3] = NumberPosition{12, Range{2, 3}}

	wantNMap := make(map[int]map[int]NumberPosition)
	wantNMap[0] = make(map[int]NumberPosition)
	wantNMap[0][1] = NumberPosition{23, Range{0, 2}}
	wantNMap[0][2] = NumberPosition{23, Range{0, 2}}

	wantNMap[2] = make(map[int]NumberPosition)
	wantNMap[2][0] = NumberPosition{1, Range{0, 0}}

	wantN := 12

	gotN := CheckAndRemoveNum(2, 2, baseNMap)

	if gotN != wantN {
		t.Errorf("Removed Wrong Number. got: %v, want: %v", gotN, wantN)
	}

	if !reflect.DeepEqual(wantNMap, baseNMap) {
		t.Errorf("Nmaps not equal. got: %v, want: %v", baseNMap, wantNMap)
	}

}

func TestPart1(t *testing.T) {
	base := "test.txt"
	want := 4361
	got, err := Part1(base)

	if err != nil {
		t.Errorf("Part 1 failed, %v", err)
	}
	if want != got {
		t.Errorf("Part 1 failed. got: %v, want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	base := "test.txt"
	want := 467835
	got, err := Part2(base)

	if err != nil {
		t.Errorf("Part 1 failed, %v", err)
	}
	if want != got {
		t.Errorf("Part 1 failed. got: %v, want: %v", got, want)
	}
}
