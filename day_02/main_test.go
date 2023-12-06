package main

import (
	"testing"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func TestFirstChallenge(t *testing.T) {
	data := utils.LoadDataSet("./data_test.txt")

	expect := 8
	got := FirstChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}

func TestSecondChallenge(t *testing.T) {
	data := utils.LoadDataSet("./data_test_2.txt")

	expect := 2286
	got := SecondChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}