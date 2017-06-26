package main

import "fmt"

func main() {
	//as an argument
	c := incrementer()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}

	// Incrementer with channel
	c1 := incrementer()
	c2 := incrementer()
	c3 := puller(c1)
	c4 := puller(c2)

	fmt.Println(<-c3, <-c4) //holding the main to exit, remember. wait to receive the baton
}

func incrementer() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
