// https://adventofcode.com/2023/day/1#part2
// 54418
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {

		line := scanner.Text()

		lineWithDigits := ReplaceStringDigits(line)

		firstDigit, lastDigit := FindDigits(lineWithDigits)

		valueStr := (string(firstDigit) + string(lastDigit))

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// fmt.Printf("%v %v %v\n", line, lineWithDigits, value)
		sum += value

	}

	fmt.Println(sum)
}

func ReplaceStringDigits(line string) string {

	var digits = map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
	}

	processedCharacters := ""
	index := 0
	occurancesSoFar := 1

	for i := 0; i < len(line); i++ {

		processedCharacters += string(line[i])

		for digit, digitName := range digits {

			if strings.HasSuffix(processedCharacters, digitName) {
				index = i - len(digitName) + occurancesSoFar
				processedCharacters = processedCharacters[:index] + digit + processedCharacters[index:]
				occurancesSoFar += 1
				break
			}
		}
	}

	return processedCharacters
}

func FindDigits(
	line string,
) (
	firstDigit rune,
	lastDigit rune,
) {

	for _, rn := range line {

		if firstDigit == 0 && unicode.IsDigit(rn) {
			firstDigit = rn
		}

		if unicode.IsDigit(rn) {
			lastDigit = rn
		}
	}

	return
}
