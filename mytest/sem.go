package main

import (
	"fmt"
	"sync"
	"time"
)

type Request struct {
	value int
}

var sem = make(chan int, 1)

type Counter struct {
	mu    sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.count
}

var counter = NewCounter()

func handle(wg *sync.WaitGroup, r Request) {
	defer wg.Done()
	//sem <- 1
	//fmt.Printf("in goroutine : %d\n", r.value)
	counter.Add()
	<-sem
}

func main() {
	wg := sync.WaitGroup{}

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		v := Request{i}
		wg.Add(1)
		//fmt.Printf("before goroutine : %d\n", i)
		sem <- 1
		go handle(&wg, v)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("---> %d : %v\n", counter.Value(), elapsed)
}
