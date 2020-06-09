package service

import "fmt"

// Sum returns the sum of the integers in the matrix
func Sum(matrix [][]int64) string {
	var total int64
	total = 0

	for _, i := range matrix {
		for _, j := range i {
			total = total + j
		}
	}

	return fmt.Sprintf("%v", total)
}
