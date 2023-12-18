// https://adventofcode.com/2023/day/2
// 1853

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	bagCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var invalidGameIds []int
	allGameIdsSum := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		gameIdPart, gamePart, found := strings.Cut(line, ": ")
		if !found {
			fmt.Println("Failed to get game id part")
		}

		_, gameIdStr, found := strings.Cut(gameIdPart, " ")
		if !found {
			fmt.Println("Failed to get game id str")
		}

		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			fmt.Println("Failed to get game id")
			os.Exit(1)
		}

		games := strings.Split(gamePart, "; ")

		for _, game := range games {

			cubeParts := strings.Split(game, ", ")

			for _, cubePart := range cubeParts {

				count, color, found := strings.Cut(cubePart, " ")
				if !found {
					fmt.Println("Failed to get cube details")
					os.Exit(1)
				}

				n, err := strconv.Atoi(count)
				if err != nil {
					fmt.Println("Failed to get cube count")
					os.Exit(1)
				}

				if n > bagCubes[color] && !slices.Contains(invalidGameIds, gameId) {
					invalidGameIds = append(invalidGameIds, gameId)
				}
			}
		}

		allGameIdsSum += gameId
	}

	invalidGameIdsSum := 0

	for _, invalidGameId := range invalidGameIds {
		invalidGameIdsSum += invalidGameId
	}

	fmt.Println(allGameIdsSum - invalidGameIdsSum)
}
