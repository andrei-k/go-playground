# Sort Integers with Goroutines

A program that sorts an array of integers by partitioning it into 4 parts of approximately equal size, and sorting each in a different goroutine. Then the main goroutine merges the 4 sorted subarrays into one sorted array. 

Each goroutine which sorts ¼ of the array prints the subarray that it will sort. When sorting is complete, the main goroutine prints the entire sorted list.

## Usage

Run in Terminal

```go
go run .     
```