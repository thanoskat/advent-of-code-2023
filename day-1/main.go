package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {

		line := scanner.Text()

		firstDigit, lastDigit := FindDigits(line)

		valueStr := (string(firstDigit) + string(lastDigit))

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		sum += value

	}

	fmt.Println(sum)
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
