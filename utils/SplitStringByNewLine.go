package utils

import "strings"


func SplitStringByNewLine(data *string) []string {
	return strings.Split(*data, "\n")
}