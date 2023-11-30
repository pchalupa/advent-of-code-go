package main

import (
	"github.com/pchalupa/advent-of-code-go/utils"
)

func main() {
	data := utils.LoadDataSet("./data.txt")

	firstChallenge := FirstChallenge(&data)
	secondChallenge := 0;

	utils.PrintResults(firstChallenge, secondChallenge)
}

func FirstChallenge(data *string) int {
	var calories []int

	elves := utils.SplitStringByEmptyLine(data)

	for _, elf := range elves {
		items := utils.SplitStringByNewLine(&elf)
		values := utils.StringsToInts(&items)
		calories = append(calories, utils.Sum(&values))
	}

	return utils.Max(&calories)
}
