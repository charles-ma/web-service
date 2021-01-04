package mymath

import (
	"fmt"
	"math/rand"
)

// PrintNum is a good method
func PrintNum(x int) {
	fmt.Println("this is a method in another package")
	fmt.Println(x)
}

// GenerateRandomNumber ...
func GenerateRandomNumber(limit int) int {
	return rand.Intn(limit) + 1
}
