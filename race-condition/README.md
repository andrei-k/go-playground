# Goroutine Race Condition

This program demonstrates an example of a **race condition**. A race condition occurs when two or more goroutines need to access the same resource. Because interleaving of goroutines is not deterministic, the value of the variable is not guaranteed to follow an expected sequence.

Here's an example of a race condition example:

| Task 1       | Task 2     |
| ------------ | ---------- |
| 1: x = 1     |            |
|              | 1: print x |
| 2: x = x + 1 |            |

| Task 1       | Task 2     |
| ------------ | ---------- |
| 1: x = 1     |            |
| 2: x = x + 1 |            |
|              | 1: print x |


Races occur due to communication between tasks. Task 1 and task 2 are communicating through the shared variable called `x`. Communication is the source for race condition.

You usually don't want to write programs where the result is not deterministic.

## Usage

Run in Terminal

```go
go run .     
```
