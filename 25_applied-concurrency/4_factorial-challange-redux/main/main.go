package main

import (
	"fmt"
	"math/rand"
)

/*
func main(){
	f := factorial(4)
	fmt.Println("Total :",<-f)
}

func factorial(n int) chan int{
	out := make(chan int)
	total := 1
	go func(){
		for i := n; i> 0;i--{
			total *= i
		}
		out <- total
	}()
	return out
}

Change the code above to execute 100 factorial computations concurrently and in parallel.
Use the pipeline pattern to accomplish this

*/

func main() {
	numbers := input(generate()...)
	c := factorial(numbers)
	for n := range c {
		fmt.Println(n)
	}
}

func generate() []int {

	var numbers []int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, 3+rand.Int()%(11-3)+1) //generate 3-15
	}
	return numbers
}

func input(numbers ...int) chan int {
	out := make(chan int)
	go func() {
		for _, val := range numbers {
			out <- val
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
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
