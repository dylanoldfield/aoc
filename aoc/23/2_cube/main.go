package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1, err := Part1("part1.txt")

	if err != nil {
		fmt.Printf("failed in part 1")
	}

	fmt.Printf("Part 1 Answer:%v \n", part1)

	part2, err := Part2("part1.txt")

	if err != nil {
		fmt.Printf("failed in part2")
	}

	fmt.Printf("Part 2 Answer: %v \n", part2)
}

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, fmt.Errorf("issue opening the file: %v", err)
	}

	var contains [3]BlockPair
	contains[0] = BlockPair{12, "red"}
	contains[1] = BlockPair{13, "green"}
	contains[2] = BlockPair{14, "blue"}

	scanner := bufio.NewScanner(file)
	roundSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		roundNum, remainder, err := GetGameNumber(line)

		if err != nil {
			return roundSum, err
		}

		maxDraws, err := GetMaxDraws(remainder)

		if err != nil {
			return roundSum, err
		}

		fits := true
		for _, contain := range contains {
			m, exists := maxDraws[contain.Colour]

			if exists && m > contain.Count {
				fits = false
				break
			}
		}

		if fits {
			roundSum += roundNum
		}
	}
	return roundSum, nil
}

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, fmt.Errorf("issue opening the file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	roundSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		_, remainder, err := GetGameNumber(line)

		if err != nil {
			return -1, err
		}

		maxDraws, err := GetMaxDraws(remainder)

		if err != nil {

			return -1, err
		}

		power := 1

		for _, count := range maxDraws {
			power *= count
		}

		roundSum += power
	}

	return roundSum, nil
}

func GetGameNumber(record string) (int, string, error) {

	gameStr := strings.Split(record, ":")
	if len(gameStr) == 1 {
		return 0, "", fmt.Errorf("no game string found")
	}

	numStr := strings.Split(gameStr[0], "Game ")

	gameNum, err := strconv.Atoi(numStr[1])

	if err != nil {
		return 0, "", err
	}

	return gameNum, strings.TrimSpace(gameStr[1]), nil

}

type BlockPair struct {
	Count int

	Colour string
}

func GetMaxDraws(drawsStrings string) (map[string]int, error) {
	maxBlockDraw := make(map[string]int)
	draws := strings.Split(drawsStrings, ";")

	for _, draw := range draws {
		blockPairs, err := GetBlockPairs(draw)

		if err != nil {
			return nil, fmt.Errorf("failed to get pairs, %v", err)
		}

		for _, blockPair := range blockPairs {

			existing, exists := maxBlockDraw[blockPair.Colour]
			if !exists || existing < blockPair.Count {
				maxBlockDraw[blockPair.Colour] = blockPair.Count
				continue
			}
		}
	}

	return maxBlockDraw, nil
}

func GetBlockPairs(draw string) ([]BlockPair, error) {
	colourPairStrings := strings.Split(draw, ",")
	var blockPairs []BlockPair
	for _, colourPairString := range colourPairStrings {
		cps := strings.TrimSpace(colourPairString)

		parts := strings.Split(cps, " ")
		num, err := strconv.Atoi(parts[0])

		if err != nil {
			return nil, fmt.Errorf("cannot get block pair num: %v ", err)
		}
		blockPairs = append(blockPairs, BlockPair{num, parts[1]})
	}

	return blockPairs, nil
}
