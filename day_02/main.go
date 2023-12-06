package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func main() {
	data := utils.LoadDataSet("./data.txt")

	firstChallenge := FirstChallenge(&data)
	secondChallenge := SecondChallenge(&data);

	utils.PrintResults(firstChallenge, secondChallenge)
}


func FirstChallenge(data *string) int {
	lines := utils.SplitStringByNewLine(data)
	result := 0
	gameLimit := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	cubesPattern := regexp.MustCompile(`(((\d{1,2}) (blue|red|green))+)`)
	gameIdPattern := regexp.MustCompile(`(\d{1,3})`)

	for _, line := range lines {
		games := strings.Split(line, ";")
		var isGamePossible = true

		for _, game := range games {
			cubesMatched := cubesPattern.FindAllString(game, -1)

			for _, cubes := range cubesMatched {
				turn :=	strings.Split(cubes, " ")
				value, _ := strconv.Atoi(turn[0])

				if(!(value <= gameLimit[turn[1]])) {
					isGamePossible = false
				}
			}
		}

		if(isGamePossible) {
			matchedGameId := gameIdPattern.FindString(line)
			gameId, _ := strconv.Atoi(matchedGameId)
			result += gameId;
		}
	}

	return result
}

func SecondChallenge(data *string) int {
	lines := utils.SplitStringByNewLine(data)
	result := 0


	cubesPattern := regexp.MustCompile(`(((\d{1,2}) (blue|red|green))+)`)

	for _, line := range lines {
		cube := map[string]int {
			"red": 1,
			"green": 1,
			"blue": 1,
		}
		games := strings.Split(line, ";")

		for _, game := range games {
			cubesMatched := cubesPattern.FindAllString(game, -1)

			for _, cubes := range cubesMatched {
				turn :=	strings.Split(cubes, " ")
				value, _ := strconv.Atoi(turn[0])

				if(cube[turn[1]] <= value) {
					cube[turn[1]] = value
				}
			}
		}

		result += (cube["red"] * cube["green"] * cube["blue"]);
	}

	return result
}
