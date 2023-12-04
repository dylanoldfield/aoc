package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	sum, err := Part1("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part1 Sum", sum)
	sum2, err := Part2("input2.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part2 Sum", sum2)

}

func Part1(fileName string) (int, error) {

	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineRune := []rune(line)
		firstDigitStr := ""
		lastDigitStr := ""
		for i := 0; i < len(lineRune); i++ {
			if unicode.IsDigit(lineRune[i]) {
				firstDigitStr = string(lineRune[i])
				break
			}
		}

		for i := len(lineRune) - 1; i >= 0; i-- {
			if unicode.IsDigit(lineRune[i]) {
				lastDigitStr = string(lineRune[i])
				break
			}
		}
		stringOffset := firstDigitStr + lastDigitStr
		intOffset, err := strconv.Atoi(stringOffset)

		if err != nil {
			fmt.Println(err)
			return 0, err
		}

		sum += intOffset
	}

	return sum, nil
}

func Part2(fileName string) (int, error) {
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	numberRegex := regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(\d)`)
	numMap := make(map[string]int)

	numMap["one"] = 1
	numMap["two"] = 2
	numMap["three"] = 3
	numMap["four"] = 4
	numMap["five"] = 5
	numMap["six"] = 6
	numMap["seven"] = 7
	numMap["eight"] = 8
	numMap["nine"] = 9

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		var numberTokensStr []string
		for i := 0; i < len(line); i++ {
			nextNum := numberRegex.FindString(line[i:])

			if nextNum != "" {
				numberTokensStr = append(numberTokensStr, nextNum)
			}
		}

		if numberTokensStr == nil {
			return 0, fmt.Errorf("no numbers found in line: %s", line)
		}

		firstNumStr := numberTokensStr[0]
		lastNumStr := numberTokensStr[len(numberTokensStr)-1]
		firstNum := 0
		lastNum := 0

		val, exists := numMap[firstNumStr]
		if !exists {
			firstNum, _ = strconv.Atoi(firstNumStr)
		} else {
			firstNum = val
		}

		val, exists = numMap[lastNumStr]
		if !exists {
			lastNum, _ = strconv.Atoi(lastNumStr)
		} else {
			lastNum = val
		}

		combined, err := strconv.Atoi(fmt.Sprint(firstNum) + fmt.Sprint(lastNum))
		if err != nil {
			return 0, err
		}

		sum += combined
	}
	return sum, nil
}
