package main

import "fmt"

/*===============================
Case
=================================
func main(){
	f := factorial(4)
	fmt.Println("Total :",f)
}

func factorial(n int) int{
	total := 1
	for i := n; i> 0;i--{
		total *= i
	}
	return total
}
Use goroutines and channels to calculate factorial
*/

func main() {
	f := factorial(4)
	fmt.Println("Total :", <-f)
}

func factorial(n int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
	}()
	return out
}
