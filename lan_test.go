package main

import (
	"fmt"
	"testing"
)

// Test files in the form of xxx_test.go
// Test methods in the form of TestX[Xx]
func TestLan(t *testing.T) {
	a, b := 1, 2
	fmt.Println(a)
	fmt.Println(b)

	m1, _ := myFunc(a, b)
	_, m2 := myFunc(a, b)

	if m1 != 3 || m2 != -1 {
		t.Error("Tests failed")
	}
}

func myFunc(a, b int) (int, int) {
	return a + b, a - b
}
