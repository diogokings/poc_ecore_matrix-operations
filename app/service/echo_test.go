package service

import "testing"

func TestEcho(t *testing.T) {

	// Test matrix 0
	callEcho(t, [][]int64{{0}}, "0")

	// Test 2x2 matrix
	callEcho(t, [][]int64{{1, 2}, {2, 3}}, "1,2\n2,3")

}

func callEcho(t *testing.T, matrix [][]int64, expected string) {
	result := Echo(matrix)

	if result != expected {
		t.Errorf("%v Failed!!!\n", t.Name())
		t.Errorf("Function echo(), with parameter %v\n", matrix)
		t.Errorf("Expected:\n%v\nbut got:\n%v\n\n", expected, result)
	}
}
