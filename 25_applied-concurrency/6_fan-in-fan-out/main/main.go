package main

import (
	"fmt"
	"sync"
)

func main() {
	//Set up the pipeline
	in := gen(2, 3)

	//Fan out
	//distribute the square across 2 goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	//Fan in
	//consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func merge(input ...chan int) chan int { //merge the result from the channels and put it in the same channel (Fan in)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(input))

	for _, input := range input {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(input) //create a new closure for the next iteration
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(nums ...int) chan int { //variadic input
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n //put the numbers into the channel
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int { //square
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
