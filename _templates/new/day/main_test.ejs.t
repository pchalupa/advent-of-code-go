---
to: day_<%= number %>/main_test.go
---
package main

import (
	"testing"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func TestFirstChallenge(t *testing.T) {
	data := utils.LoadDataSet("./data_test.txt")
	channel := make(chan int)

	expect := 0

	go FirstChallenge(&data, channel)

	got := <- channel


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}

func TestSecondChallenge(t *testing.T) {
	data := utils.LoadDataSet("./data_test_2.txt")
	channel := make(chan int)

	expect := 0

	go SecondChallenge(&data, channel)

	got := <- channel


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}