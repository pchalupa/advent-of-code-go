package utils

import "strings"


func SplitStringByEmptyLine(data *string) []string {
	return strings.Split(*data, "\n\n")
}