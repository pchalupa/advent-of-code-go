package main

import (
	"regexp"
	"strconv"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func main() {
	data := utils.LoadDataSet("./data.txt")

	firstChallenge := FirstChallenge(&data)
	secondChallenge := SecondChallenge(&data);


	utils.PrintResults(firstChallenge, secondChallenge)
}

func FirstChallenge(data *string) int {
	sum := 0;
	lines := utils.SplitStringByNewLine(data)
	pattern := regexp.MustCompile(`(\d)`)

	for _, line := range lines {
		matched := pattern.FindAllString(line, -1)

		first, last := matched[0], matched[len(matched)-1]

		value, _ := strconv.Atoi(first + last)

		sum +=value
	}

	return sum
}

func SecondChallenge(data *string) int {
	numbers := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four":4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	sum := 0;
	lines := utils.SplitStringByNewLine(data)
	fromLeft := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	fromRight := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d)`)

	for _, line := range lines {
		first := fromLeft.FindString(line)
		last := utils.Revert(fromRight.FindString(utils.Revert(line)))


		if(numbers[first] > 0) {
			first = strconv.Itoa(numbers[first])
		}


		if(numbers[last] > 0) {
			last = strconv.Itoa(numbers[last])
		}

		value, _ := strconv.Atoi(first + last)

		sum += value
	}

	return sum
}
