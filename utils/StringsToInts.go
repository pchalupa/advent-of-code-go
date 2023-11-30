package utils

import "strconv"

func StringsToInts(items * []string) []int {
	var values []int;

	for _, item := range *items {
		value, err := strconv.Atoi(item)

		if (err != nil) {
			panic("Not able to convert string to int")
		}

		values = append(values, value)

	}

	return values
}