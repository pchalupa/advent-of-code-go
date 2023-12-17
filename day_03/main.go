package main

import (
	"fmt"
	"regexp"
	"slices"
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

type Item struct {
	Value int
	Coordinates []string
}

func FirstChallenge(data *string) int {
	lines := utils.SplitStringByNewLine(data)
	characters := make(map[string]bool)
	var items []Item;
	result := 0;

	numbersPattern := regexp.MustCompile(`(\d{1,3})`)
	charactersPattern := regexp.MustCompile(`([^\w.\n])`)

	for lineNumber, line := range lines {
		valueMatches := numbersPattern.FindAllString(line, -1)
		valueLength := numbersPattern.FindAllStringIndex(line, -1)
		charactersMatches := charactersPattern.FindAllStringIndex(line, -1)

		for i, match := range valueMatches {
			var coordinates []string

			for valueLength[i][0] < valueLength[i][1] {
				coordinates = append(coordinates, fmt.Sprintf("%d,%d", valueLength[i][0], lineNumber))
				valueLength[i][0]++
			}
			value,_ := strconv.Atoi(match);
			items = append(items, Item{Value: value, Coordinates: coordinates})
		}

		for _, match := range charactersMatches {
			coordinates := fmt.Sprintf("%d,%d", match[0], lineNumber)
			characters[coordinates] = true
		}
	}

	for _, item := range items {
		for _, coordinate := range item.Coordinates {
			foo := strings.Split(coordinate, ",");
			x,_ := strconv.Atoi(foo[0])
			y,_ := strconv.Atoi(foo[1])


			for i := x-1; i <= x+1; i++ {
				for j := y-1; j <= y+1; j++ {
					if(i < 0 || j < 0) {
						 continue;
					}
					testCoordinates := fmt.Sprintf("%d,%d", i, j);

					if(characters[testCoordinates] && item.Value > 0) {
						result = result + item.Value
						item.Value = 0
					}
				}
			}
		}
	}

	return result
}

func SecondChallenge(data *string) int {
	lines := utils.SplitStringByNewLine(data)
	characters := make(map[string]string)
	result := 0;
	numbers := make(map[string]int)

	numbersPattern := regexp.MustCompile(`(\d{1,3})`)
	charactersPattern := regexp.MustCompile(`([^\w.\n])`)

	for lineNumber, line := range lines {
		valueMatches := numbersPattern.FindAllString(line, -1)
		valueLength := numbersPattern.FindAllStringIndex(line, -1)
		charactersLength := charactersPattern.FindAllStringIndex(line, -1)

		// values
		for i, match := range valueMatches {
			value,_ := strconv.Atoi(match);

			for valueLength[i][0] < valueLength[i][1] {
				numbers[fmt.Sprintf("%d,%d", valueLength[i][0], lineNumber)] = value
				valueLength[i][0]++
			}
		}

		// characters
		for _, match := range charactersLength {

			coordinates := fmt.Sprintf("%d,%d", match[0], lineNumber)
			characters[coordinates] = coordinates;
		}
	}

	for _, character := range characters  {
		coordinates := strings.Split(character, ",");
		x,_ := strconv.Atoi(coordinates[0])
		y,_ := strconv.Atoi(coordinates[1])

		var ratio []int;

		for i := x-1; i <= x+1; i++ {
			for j := y-1; j <= y+1; j++ {
				value, ok:=numbers[fmt.Sprintf("%d,%d", i, j)]

				if(ok) {
					if(!slices.Contains(ratio, value)) {
						ratio = append(ratio, value)
					}
				}
			}
		}
		if(len(ratio) >= 2) {
			result += ratio[0] * ratio[1]
		}

	}



	return result
}
