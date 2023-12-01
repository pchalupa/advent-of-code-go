---
to: day_<%= number %>/main_test.go
---
package main

import (
	"testing"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func setup() string {
	return utils.LoadDataSet("./data_test.txt")
}

func TestFirstChallenge(t *testing.T) {
	data := setup()

	expect := 0
	got := FirstChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}

func TestSecondChallenge(t *testing.T) {
	data := setup()

	expect := 0
	got := SecondChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}