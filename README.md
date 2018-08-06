# goPra
This is a Golang Playground. 

Goal: get familiar with Golang syntax and able to solve some simple problems using this language.

## Golang Notes
### Go characteristics
1. Go is a compiled language(rather than scripting language, e.g Python), compiled to machine code. Code may need to be compiled separately for different computer architecture. Scripting languages are often cross platform
2. Built-in functions are organized in modules. e.g fmt, time

### Commands
1. "go run" will compile and run your code in a temporary work space, you can use "go run --work" to see the dir
2. "go build" to compile the code
3. Program entry point is "main" func inside "main" package, similar to Java
4. godoc -http=:(portNum) to start local doc
5. "go get" can fetch 3rd party libraries from various sources such as github(internet). If this command executed within workspace, go will search all the source files for 3rd party packages and fetch them. Our source files become like package.json(node package config file)
6. There are other go package mgmt tool available, such as [goop](https://github.com/petejkim/goop)
7. "go fmt" can format your code

### Types
1. int8/16/32/64, uint8/16/32/64, byte = uint8, rune = int32
2. machine(computer architecture) dependent: int, uint, uintptr
3. float32/64
4. complex64/128
5. NaN
6. Inf/-Inf
7. strings supports "" and `` (different from js template literals, cannot take ${})


### Syntax
1. Go is C like. Java/C/C++/Js are also C like
2. Go gets rid of () around conditions. Also gets rid of ; as line breaker
3. null is nil
4. if statement can have var assignment
5. type conversion: a.(int), to get type a.(type), but only applied to empty interface type. a.(type) can only be used within a type switch statement

### Variable and Declaration
1. := can infer type to var. declare and assign value. can be used as long as one var is new
2. vars cannot be declared multiple times
3. won't compile if declared var not used
4. declare var: var v type / assign var v = value / declare and assign v := value
5. var will be initiated to zero value if not assigned any. zero value meaning 0, false, ""
6. v := value can only be used within functions

### Package
1. won't compile if imported package not used
2. Packages are imported using relative path to $GOPATH
3. Source file structure: package declaration -> imports -> logic, same as Python
4. Name starting with caps will be exported from a module

### Struct
1. No class, only struct
2. No constructors, only customized factory methods
3. new(X) is equal to &X{}, will only allocate memory for the new variable
4. class method are implemented using method receivers

### Array & slice
1. Array length is fixed!!!
2. Use slice more

### Test
1. run "go test"
2. test cases are searched by file name --- xxx_test.go
3. test methods are searched by name --- TestX[Xx]

### Visibility
1. Visibility is based on the case of names
2. Types, functions and struct fields are accessible to outside of a package when name starts with uppercase letter. Otherwise only accessible to package

### Interface
1. just a contract
2. everything implements the empty interface

### Exception
1. Go does have panic(throw) and recover(catch) clauses, but are rarely used. Normally should throw an exception using multiple return values
2. Defer clause is like finally, will defer operation to when enclosing function returns(can have multiple return points)

### Function type
1. type Add func(a int, b int) int

### Concurrency
1. go routine is invoked using go myFunction()
2. go routine will start go thread, having m:n mapping with os threads

## Books
[The little Go book](https://www.openmymind.net/assets/go/go.pdf) *Free*

