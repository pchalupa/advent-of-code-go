---
to: day_<%= number %>/test_main.go
---
package main

import (
	"testing"

	"github.com/pchalupa/advent-of-code-go/utils"
)

func setup() string {
	return utils.LoadDataSet("./test_data.txt")
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
	got := FirstChallenge(&data)


	if got != expect {
		t.Errorf("Expect %d, but got %d", expect, got)
	}
}