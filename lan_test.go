package main

import (
	"fmt"
	"strconv"
	"testing"
)

var topicCount int

// Test files in the form of xxx_test.go
// Test methods in the form of TestX[Xx]
// Test untilities
func countTopic() {
	topicCount++
	fmt.Println("\nTopic: " + strconv.Itoa(topicCount))
}

// Variable
func myFunc(a, b int) (int, int) {
	return a + b, a - b
}

func TestVar(t *testing.T) {
	countTopic()
	a, b := 1, 2
	fmt.Println(a)
	fmt.Println(b)

	m1, _ := myFunc(a, b)
	_, m2 := myFunc(a, b)

	if m1 != 3 || m2 != -1 {
		t.Error("Tests failed")
	}
}

// Struct
type MyClass struct {
	id   int
	Name string
}

func (instance *MyClass) myPrintOnStruct() {
	fmt.Println("id: " + strconv.Itoa(instance.id) + " Name: " + instance.Name)
}

func TestStruct(t *testing.T) {
	countTopic()

	myInstance := &MyClass{
		id:   5,
		Name: "Mike",
	}

	myInstance.myPrintOnStruct()
}
