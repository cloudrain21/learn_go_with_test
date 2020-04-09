package main

import (
	"fmt"
	"sync"
)

const numWorker = 4

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

type Result struct {
	who int
	num int
	ret int
}

var wg sync.WaitGroup

func main() {
	jobs := make(chan int, 100)
	results := make(chan Result, 100)

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, results)
	}

	for i := 1; i <= 40; i++ {
		jobs <- i
	}

	for r := range results {
		fmt.Printf("%d : %d : %d\n", r.who, r.num, r.ret)
	}

	wg.Wait()
}

func worker(who int, jobs <-chan int, results chan<- Result) {
	defer wg.Done()
	for n := range jobs {
		results <- Result{who, n, fib(n)}
	}
}
