# Benchmarking in Go

*Work in progress*

Simple program to practice writing benchmarking tests following this [article](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/).  

Benchmark tests collect a measurement by executing a function multiple times. The `b.N` value is increased  on each execution inside a loop.  

## Usage

Run in Terminal

Simply run the program without compiling it.
```go
go run .
```

Run all tests using verbose mode.
```go
go test -v
```

Run all benchmarks in the current project. The `bench` flag filters benchmark function names with a regular expression. Use `-bench=.` to run all benchmarks.
```go
go test -bench=.
```

Run benchmarks multiple times by using a `-count` flag.
```go
go test -bench=. -count 5
```