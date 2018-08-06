package main

import "testing"

func TestTwoSum(t *testing.T) {
	a := []int{3, 9, 12, 6, 7}
	target := 9
	r := TwoSum(a, target)
	if (r[0] != 0 || r[1] != 3) && (r[0] != 3 || r[1] != 0) {
		t.Errorf("test failed")
	}
}
