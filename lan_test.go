package main

import (
	"fmt"
	"testing"
)

// Test files in the form of xxx_test.go
// Test methods in the form of TestX[Xx]
func TestLan(t *testing.T) {
	a, b := 1, "nice"
	fmt.Println(a)
	fmt.Println(b)
}
