package service

import "testing"

func TestSum(t *testing.T) {

	// Test matrix 0
	callSum(t, [][]int64{{0}}, "0")

	// Test 2x2 matrix
	callSum(t, [][]int64{{1, 4}, {2, 3}}, "10")

	// Test 3x3 matrix
	callSum(t, [][]int64{{100, 1, 2}, {0, 3, 20}, {4, 5, 5}}, "140")

}

func callSum(t *testing.T, matrix [][]int64, expected string) {
	result := Sum(matrix)

	if result != expected {
		t.Errorf("%v Failed!!!\n", t.Name())
		t.Errorf("Function sum(), with parameter %v\n", matrix)
		t.Errorf("Expected:\n%v\nbut got:\n%v\n\n", expected, result)
	}
}
