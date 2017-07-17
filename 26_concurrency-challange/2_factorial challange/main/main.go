package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Use fan in fan out to calculate 1000 factorial running consecutively

*/

func main() {
	numbers := generate()

	//Fan Out
	//distribute to 10 goroutines
	c1 := factorial(numbers)
	c2 := factorial(numbers)
	c3 := factorial(numbers)
	c4 := factorial(numbers)
	c5 := factorial(numbers)
	c6 := factorial(numbers)
	c7 := factorial(numbers)
	c8 := factorial(numbers)
	c9 := factorial(numbers)
	c10 := factorial(numbers)

	var index int
	for n := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		index++
		fmt.Println(index, "		:	", n)
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

func generate() <-chan int {

	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			//generate 3-15
			out <- 3 + rand.Int()%(11-3) + 1
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
