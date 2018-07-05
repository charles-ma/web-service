package mymath

import "fmt"

// PrintNum is a good method
func PrintNum(x int) {
	fmt.Println("this is a method in another package")
	fmt.Println(x)
}
