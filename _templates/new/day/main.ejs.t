---
to: day_<%= number %>/main.go
---
package main

import (
	"github.com/pchalupa/advent-of-code-go/utils"
)

func main() {
	data := utils.LoadDataSet("./data.txt")

	firstChallenge := FirstChallenge(&data)
	secondChallenge := SecondChallenge(&data);

	utils.PrintResults(firstChallenge, secondChallenge)
}

func FirstChallenge(data *string) int {
	return 0
}

func SecondChallenge(data *string) int {
	return 0
}
