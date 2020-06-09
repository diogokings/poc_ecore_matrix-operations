package service

import "fmt"

// Flatten returns the matrix as one line string, with values separated by commas
func Flatten(matrix [][]int64) string {
	var str = ""

	for _, i := range matrix {
		for _, j := range i {
			var n = fmt.Sprintf("%v", j)
			str = str + n + ","
		}
	}

	size := len(str)

	if size > 0 {
		str = str[:size-1]
	}

	return str
}
