package utils

import "os"

func LoadDataSet(path string) string {
	data, err := os.ReadFile(path)

		if err != nil {
		panic("Not able to read")
	}

	return string(data)
}