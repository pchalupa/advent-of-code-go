package utils

// FInds the max int int the slice
func Max(values * []int) int {
	result := 0;

	for _, value := range *values {
		if value > result {
			result = value
		}
	}

	return result
}