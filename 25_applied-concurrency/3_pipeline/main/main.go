package main

import "fmt"

/*===========================================
Pipeline
===========================================
Series of stages connected by channels
*/

func main() {
	//Set up the pipeline
	c := gen(2, 3)
	out := sq(c)

	// consume the ouput.
	fmt.Println(<-out) //4
	fmt.Println(<-out) //9

	//Refactor
	for n := range sq(sq(gen(2, 3))) { //first call the gen function,
		// that returns channel then pass it to square,
		// return channel andpass it again to the sq function
		fmt.Println(n)
	}
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
