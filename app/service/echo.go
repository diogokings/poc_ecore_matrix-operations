package service

import "fmt"

// Echo will return the matrix, informed in the paramenter, as a string, in matrix format
func Echo(matrix [][]int64) string {
	var strMatrix = ""

	for _, i := range matrix {
		var line = ""

		if len(strMatrix) > 0 {
			strMatrix = strMatrix + "\n"
		}

		for _, j := range i {
			var n = fmt.Sprintf("%v", j)
			line = line + n + ","
		}

		size := len(line)

		if size > 0 {
			line = line[:size-1]
		}

		strMatrix = strMatrix + line
	}

	return strMatrix
}
