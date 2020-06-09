package service

import "fmt"

// Multiply returns the product of the integers in the matrix
func Multiply(matrix [][]int64) string {
	var total int64
	total = 1

	for _, i := range matrix {
		for _, j := range i {
			total = total * j
		}
	}

	return fmt.Sprintf("%v", total)
}
