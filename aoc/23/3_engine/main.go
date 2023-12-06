package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file1 := "input1.txt"

	part1, err := Part1(file1)

	if err != nil {
		fmt.Printf("part1 failure. %v", err)
	}

	fmt.Printf("Part 1 Answer: %v", part1)

}

type NumberPosition struct {
	Value int
	Range Range
}

type Range struct {
	Start int
	End   int
}

type Boundary struct {
	row    int
	column int
}

func Part1(filename string) (int, error) {
	sMap, nMap, boundary, err := LoadSchematic(filename)

	if err != nil {
		return -1, fmt.Errorf("failed to load schematic %v", err)
	}

	sum := 0
	for y, val := range sMap {
		for _, x := range val {
			// up
			if y > 0 {
				sum += CheckAndRemoveNum(y-1, x, nMap)
			}

			// down
			if y < boundary.row {
				sum += CheckAndRemoveNum(y+1, x, nMap)
			}

			// left
			if x > 0 {
				sum += CheckAndRemoveNum(y, x-1, nMap)
			}

			// right
			if x < boundary.column {
				sum += CheckAndRemoveNum(y, x+1, nMap)
			}

			// top-left
			if y > 0 && x > 0 {
				sum += CheckAndRemoveNum(y-1, x-1, nMap)
			}

			// top-right
			if y > 0 && x < boundary.column {
				sum += CheckAndRemoveNum(y-1, x+1, nMap)
			}

			// bottom-left
			if y < boundary.row && x > 0 {
				sum += CheckAndRemoveNum(y+1, x-1, nMap)
			}

			// bottom-right
			if y < boundary.row && x < boundary.column {
				sum += CheckAndRemoveNum(y+1, x+1, nMap)
			}

		}
	}

	return sum, nil
}

func CheckAndRemoveNum(y int, x int, nMap map[int]map[int]NumberPosition) int {
	row, exists := nMap[y]

	if !exists {
		return 0
	}

	num, exists := row[x]

	if !exists {
		return 0
	}

	for i := num.Range.Start; i <= num.Range.End; i++ {
		delete(nMap[y], i)
	}

	return num.Value
}

func addToNumMap(curNum string, row int, start int, end int, numMap map[int]map[int]NumberPosition) (map[int]map[int]NumberPosition, error) {
	if len(curNum) > 0 {
		num, err := strconv.Atoi(curNum)
		if err != nil {
			return numMap, fmt.Errorf("unable to convert num, %v", err)
		}

		if _, exists := numMap[row]; !exists {
			numMap[row] = make(map[int]NumberPosition)
		}

		for i := start; i < end+1; i++ {
			numMap[row][i] = NumberPosition{num, Range{start, end}}
		}
	}

	return numMap, nil
}

func LoadSchematic(filename string) (map[int][]int, map[int]map[int]NumberPosition, Boundary, error) {
	file, err := os.Open(filename)
	numMap := make(map[int]map[int]NumberPosition)
	symbolMap := make(map[int][]int)
	if err != nil {
		return symbolMap, nil, Boundary{}, err
	}

	scanner := bufio.NewScanner(file)
	y := 0
	lineWidth := 0
	for scanner.Scan() {
		line := scanner.Text()
		curNum := ""
		start := -1
		lineWidth = len(line)

		// iterate through the rows add number to map for every x,y coordinate they are apart of

		var symbolXs []int
		for x, r := range line {
			// check if rune is a number
			if unicode.IsDigit(r) {
				curNum += string(r)
				// check if first number we have seen since "." or symbol
				if start == -1 {
					start = x
				}
			} else {
				numMap, err = addToNumMap(curNum, y, start, x-1, numMap)
				if err != nil {
					return symbolMap, nil, Boundary{}, err
				}
				curNum = ""
				start = -1

				if r != '.' {
					symbolXs = append(symbolXs, x)
				}
			}

		}
		numMap, err = addToNumMap(curNum, y, start, len(line)-1, numMap)
		if err != nil {
			return symbolMap, nil, Boundary{}, err
		}
		if len(symbolXs) > 0 {
			symbolMap[y] = symbolXs
		}
		y++
	}

	return symbolMap, numMap, Boundary{y - 1, lineWidth - 1}, nil
}
