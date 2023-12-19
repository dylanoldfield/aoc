package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p1, err := Part1("part1.txt")

	if err != nil {
		fmt.Printf("Part 1 failed with error: %v", err)
	}

	fmt.Printf("Part 1 Score: %v\n", p1)

	p2, err := Part2("part2.txt")

	if err != nil {
		fmt.Printf("Part 2 failed with error: %v", err)
	}

	fmt.Printf("Part 2 Score: %v\n", p2)

}

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(file)
	totalScore := 0
	cardNum := 0

	for scanner.Scan() {
		cardNum++
		line := scanner.Text()
		numbers := strings.Split(line, ":")

		if len(numbers) != 2 {
			return -1, fmt.Errorf("line format doesn't match Card#: #### | ####")
		}
		scratchSets := strings.Split(strings.TrimSpace(numbers[1]), "|")

		if len(scratchSets) != 2 {
			return -1, fmt.Errorf("could not find two sets split by |")
		}

		winningStrings := strings.Split(strings.TrimSpace(scratchSets[0]), " ")
		fmt.Printf("winning strings: %v\n", winningStrings)
		scratchStrings := strings.Split(strings.TrimSpace(scratchSets[1]), " ")

		fmt.Printf("scratch strings: %v\n", scratchStrings)
		winningSet := make(map[string]bool)

		for _, item := range winningStrings {
			winningSet[item] = true
		}

		var matched []string
		score := 0

		for _, item := range scratchStrings {
			_, exists := winningSet[item]

			if exists && item != "" {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}

				matched = append(matched, item)
			}

			delete(winningSet, item)
		}

		fmt.Printf("Card %v score: %v\n", cardNum, score)
		fmt.Printf("Matched: %v\n", matched)
		totalScore += score
	}

	return totalScore, nil
}

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(file)
	cardNum := 0
	var scores []int

	for scanner.Scan() {
		cardNum++
		line := scanner.Text()
		numbers := strings.Split(line, ":")

		if len(numbers) != 2 {
			return -1, fmt.Errorf("line format doesn't match Card#: #### | ####")
		}
		scratchSets := strings.Split(strings.TrimSpace(numbers[1]), "|")

		if len(scratchSets) != 2 {
			return -1, fmt.Errorf("could not find two sets split by |")
		}

		winningStrings := strings.Split(strings.TrimSpace(scratchSets[0]), " ")
		scratchStrings := strings.Split(strings.TrimSpace(scratchSets[1]), " ")

		winningSet := make(map[string]bool)

		for _, item := range winningStrings {
			winningSet[item] = true
		}

		score := 0

		for _, item := range scratchStrings {
			_, exists := winningSet[item]

			if exists && item != "" {
				score++
			}

			delete(winningSet, item)
		}

		scores = append(scores, score)
	}

	total := 0
	totals := make([]int, len(scores))
	for i := len(scores) - 1; i >= 0; i-- {

		j := i + 1

		sum := 0
		for j < len(scores) && j <= i+scores[i] {
			sum += totals[j]
			j++
		}

		ovflow := func() int {
			if scores[i]+i >= len(scores) {
				return len(scores) - 1 - i
			}
			return scores[i]
		}

		totals[i] = ovflow() + sum
		total += totals[i]
	}

	fmt.Printf("totals: %v\n", totals)

	return total + len(scores), nil
}
