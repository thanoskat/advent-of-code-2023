// https://advent-of-code-2023/day/2
// 72706

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()

		_, gamePart, found := strings.Cut(line, ": ")
		if !found {
			fmt.Println("Failed to get game id part")
		}

		// _, gameIdStr, found := strings.Cut(gameIdPart, " ")
		// if !found {
		// 	fmt.Println("Failed to get game id str")
		// }

		// gameId, err := strconv.Atoi(gameIdStr)
		// if err != nil {
		// 	fmt.Println("Failed to get game id")
		// 	os.Exit(1)
		// }

		max := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		draws := strings.Split(gamePart, "; ")

		for _, draw := range draws {

			cubeParts := strings.Split(draw, ", ")

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

				if max[color] < n {
					max[color] = n
				}
			}
		}

		gamePower := max["red"] * max["green"] * max["blue"]

		sum += gamePower
	}

	fmt.Println(sum)
}
