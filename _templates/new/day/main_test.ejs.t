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

	expect := 0
	got := FirstChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}

func TestSecondChallenge(t *testing.T) {
	data := utils.LoadDataSet("./data_test_2.txt")

	expect := 0
	got := SecondChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}