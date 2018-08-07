package main

import "testing"

func TestHammingDistance(t *testing.T) {
	a := 1
	b := 4
	if HammingDistance(a, b) != 2 {
		t.Errorf("Test failed")
	}
}
