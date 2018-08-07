package main

// not very often to see () in import statement though
import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"
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

	// a := [5]int{0, 1, 2, 3, 4}
	a := [...]int{0, 1, 2, 3, 4}
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

// Interface
func TestInterface(t *testing.T) {
	countTopic()

	v := MyInterImplementation{}
	myInterfaceTest(v)

	var value interface{}
	if value != nil {
		fmt.Println(value.(int))
	}
}

func myInterfaceTest(inter MyInterface) {
	inter.MyFunction("this is a message")
}

type MyInterImplementation struct {
}

func (imp MyInterImplementation) MyFunction(message string) {
	fmt.Println(message)
}

type MyInterface interface {
	MyFunction(message string)
}

// Concurrency/go routine
func TestConcurrency(t *testing.T) {
	countTopic()

	for i := 0; i < 10; i++ {
		go inc()
	}
	time.Sleep(time.Millisecond * 10)
}

var count int
var lock sync.Mutex

func inc() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println(count)
}

// Concurrency/channel
func TestChannel(t *testing.T) {
	countTopic()

	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	c <- 0
	time.Sleep(time.Millisecond * 10)
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("Worker %d got %d\n", w.id, data)
		if data < 10 {
			c <- (data + 1)
		}
	}
}

func TestTypes(t *testing.T) {
	countTopic()
	var f float64
	f = 32
	fmt.Println(f)

	var s string
	// `` is not taking escaped chars
	s = `good e\nxample`
	fmt.Println(s)

	if s == `good e\nxample` {
		// without using libraries
		println("print again")
		// using libraries
		fmt.Println("print formatted again")
	}
}

func TestFunctional(t *testing.T) {
	countTopic()

	type adder func(a int) int
	var addOne adder

	adderFactory := func(a int) adder {
		var myInnerFunc adder = func(b int) int {
			return a + b
		}
		return myInnerFunc
	}
	addOne = adderFactory(1)
	addTen := adderFactory(10)

	n, _ := fmt.Println(addOne(2))
	fmt.Println(n)
	n, _ = fmt.Println(addTen(999), "good")
	fmt.Println(n)
}

func TestScope(t *testing.T) {
	countTopic()

	type adder struct {
		// yet another adder
	}

	fmt.Println(os.Args)
}

func TestPointers(t *testing.T) {
	countTopic()
	i := 1
	const Pi = 3.14
	p := &i
	*p = 2000
	fmt.Println(i)
}

func TestString(t *testing.T) {
	countTopic()
	const s = "Hello, 世界"
	for i, c := range s {
		fmt.Println(i, c, string(c), reflect.TypeOf(c))
	}
	var r rune
	r = 97
	fmt.Println(string(r), reflect.TypeOf(r))
}
