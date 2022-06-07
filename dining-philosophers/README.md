# Dining Philosopher’s Problem

There are 5 philosophers sitting at a round table. 1 chopstick is placed between each adjacent pair. Each philosopher wants to eat rice from their plate, but needs 2 chopsticks. Only one philosopher can hold a chopstick at a time and there are not enough chopsticks for everyone to eat at once.

Dining philosophers is a classic problem involving concurrency and synchronization. 
https://en.wikipedia.org/wiki/Dining_philosophers_problem  

This program implement's the problem with the following constraints/modifications.  

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times.

3. The philosophers pick up the chopsticks in any order.

4. In order to eat, a philosopher must get permission from a **host** which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat concurrently.

6. Each philosopher is numbered, 1 through 5.

7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

## Usage

Run in Terminal

```go
go run .     
```
