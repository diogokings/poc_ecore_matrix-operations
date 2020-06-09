package service

import "testing"

func TestInvert(t *testing.T) {

	// Test matrix 0
	callInvert(t, [][]int64{{0}}, "0")

	// Test 2x2 matrix
	callInvert(t, [][]int64{{1, 4}, {2, 3}}, "1,2\n4,3")

	// Test 3x3 matrix
	callInvert(t, [][]int64{{99, 4, 5}, {2, 3, 1900}, {59, 5, 5}}, "99,2,59\n4,3,5\n5,1900,5")

}

func callInvert(t *testing.T, matrix [][]int64, expected string) {
	result := Invert(matrix)

	if result != expected {
		t.Errorf("%v Failed!!!\n", t.Name())
		t.Errorf("Function invert(), with parameter %v\n", matrix)
		t.Errorf("Expected:\n%v\nbut got:\n%v\n\n", expected, result)
	}
}
