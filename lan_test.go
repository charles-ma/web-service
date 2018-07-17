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
// Simple Struct
type MyClass struct {
	id   int
	Name string
}

func (instance *MyClass) myPrintOnStruct() {
	fmt.Println("id: " + strconv.Itoa(instance.id) + " Name: " + instance.Name)
}

func TestStruct(t *testing.T) {
	countTopic()

	myInstance := &MyClass{5, "Mike"}

	myInstance.myPrintOnStruct()
}

// More complicated struct
type Saiyan struct {
	*Person // this field is implicitly declared
	Father  *Saiyan
}

type Person struct {
	Name  string
	Power int
}

func TestComplicatedStruct(t *testing.T) {
	countTopic()

	mySaiyan := &Saiyan{&Person{"kakasi", 100}, &Saiyan{&Person{"K Father", 80}, nil}}

	if mySaiyan.Father != nil {
		fmt.Println(mySaiyan.Name)
	}
}

// Array
func TestArray(t *testing.T) {
	countTopic()

	a := [5]int{0, 1, 2, 3, 4}
	for index, value := range a {
		fmt.Print(a[index])
		fmt.Print(value)
	}
	fmt.Println()

	// new to allocate space to array
	b := new([6]int)
	b[0] = 1
	fmt.Println(*b)

	// slice example
	// c & d share the same backing array, but are different slices
	// c := make([]int, 0, 10)
	// initial capacity not required
	var c []int
	d := append(c, 5)
	fmt.Println(d)
	fmt.Println(c)

	// array grow test
	s := make([]int, 0, 5)
	fmt.Println(cap(s))
	for i := 0; i < 25; i++ {
		s = append(s, i)
	}
	fmt.Println(cap(s))

	// if slice length not specified, slice has init value in it
	s = make([]int, 5)
	// s = make([]int, 0, 5)
	s = append(s, 5)
	fmt.Println(s)
}

// Map
func TestMap(t *testing.T) {
	countTopic()

	m := make(map[string]int)
	m["conan"] = 100
	power, exists := m["conan"]
	fmt.Println(power, exists)
	p, e := m["kakasi"]
	fmt.Println(p, e)
}
