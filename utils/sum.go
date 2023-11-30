package utils

// Sums values in a slice
func Sum(values *[]int) int {
	var result int;

	for _, value := range *values {
		result += value
	}

	return result
}
