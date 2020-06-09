package service

// Invert returns the matrix as a string in matrix format, where the columns and rows are inverted
func Invert(matrix [][]int64) string {
	var xLength = len(matrix)
	var yLength = len(matrix[0])

	inverted := make([][]int64, yLength)

	for i := range inverted {
		inverted[i] = make([]int64, xLength)
	}

	for i := 0; i < xLength; i++ {
		for j := 0; j < yLength; j++ {
			inverted[j][i] = matrix[i][j]
		}
	}
	var str = Echo(inverted)
	return str
}
