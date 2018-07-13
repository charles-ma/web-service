# goPra
This is a Golang Playground. 

Goal: get familiar with Golang syntax and able to solve some simple problems using this language.

## Golang Notes
### Go characteristics
1. Go is a compiled language, compiled to assembly

### Commands
1. "go run" will compile and run your code in a temporary work space, you can use "go run --work" to see the dir
2. "go build" to compile the code
3. Program entry point is "main" func inside "main" package
4. godoc -http=:(portNum) to start local doc

### Syntax
1. Go is C like. Java/C/C++/Js are also C like
2. Go gets rid of () around conditions. Also gets rid of ; as line breaker

### Variable and Declaration
1. := can infer type to var. declare and assign value. can be used as long as one var is new
2. vars cannot be declared multiple times
3. won't compile if declared var not used
4. declare var: var v type / assign var v = value / declare and assign v := value
5. var will be initiated to zero value if not assigned any. zero value meaning 0, false, ""

### Package
1. won't compile if imported package not used

### Test
1. run "go test"
2. test cases are searched by file name --- xxx_test.go
3. test methods are searched by name --- TestX[Xx]

