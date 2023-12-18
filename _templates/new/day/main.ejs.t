---
to: day_<%= number %>/main.go
---
package main

import (
	"github.com/pchalupa/advent-of-code-go/utils"
)

func main() {
	data := utils.LoadDataSet("./data.txt")
	firstChallenge := make(chan int)
	secondChallenge := make(chan int)

	go FirstChallenge(&data, firstChallenge)
	go SecondChallenge(&data, secondChallenge);

	utils.PrintResults(<-firstChallenge, <-secondChallenge)
}

func FirstChallenge(data *string, channel chan int) {
	var result = 0

	channel <- result
}

func SecondChallenge(data *string, channel chan int) {
	var result = 0

	channel <- result
}
