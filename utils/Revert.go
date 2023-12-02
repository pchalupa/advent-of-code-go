package utils

// Reverts a string
func Revert(value string) string {
	result := ""

	for _, letter:=range value {
		result = string(letter) + result;
	}

	return result
}