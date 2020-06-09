package service

import "testing"

func TestFlatten(t *testing.T) {

	// Test matrix 0
	callFlatten(t, [][]int64{{0}}, "0")

	// Test 2x2 matrix
	callFlatten(t, [][]int64{{1, 2}, {2, 3}}, "1,2,2,3")

}

func callFlatten(t *testing.T, matrix [][]int64, expected string) {
	result := Flatten(matrix)

	if result != expected {
		t.Errorf("%v Failed!!!\n", t.Name())
		t.Errorf("Function flatten(), with parameter %v\n", matrix)
		t.Errorf("Expected:\n%v\nbut got:\n%v\n\n", expected, result)
	}
}
