package service

import "testing"

func TestMultiply(t *testing.T) {

	// Test matrix 0
	callMultiply(t, [][]int64{{0}}, "0")

	// Test 2x2 matrix
	callMultiply(t, [][]int64{{1, 4}, {2, 5}}, "40")

	// Test 3x3 matrix with 0
	callMultiply(t, [][]int64{{100, 1, 2}, {0, 3, 20}, {4, 5, 5}}, "0")

	// Test 3x3 matrix
	callMultiply(t, [][]int64{{100, 1, 2}, {1, 3, 1}, {1, 2, 1}}, "1200")

}

func callMultiply(t *testing.T, matrix [][]int64, expected string) {
	result := Multiply(matrix)

	if result != expected {
		t.Errorf("%v Failed!!!\n", t.Name())
		t.Errorf("Function multiply(), with parameter %v\n", matrix)
		t.Errorf("Expected:\n%v\nbut got:\n%v\n\n", expected, result)
	}
}
