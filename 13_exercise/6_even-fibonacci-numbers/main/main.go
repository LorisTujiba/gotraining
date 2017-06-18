package main

import "fmt"

/*
	Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2,
the first 10 terms will be:
							1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the
even-valued terms.*/

func fibonacci(input int) int{

	var lh = 0
	var rh = 1
	var eh int
	var even int


	for i:=0;i<input;i++{
		eh = lh+rh
		lh=rh
		rh=eh
		if eh%2==0{
			even+=eh
		}
		if eh > 4000000{
			eh = lh
			break
		}
	}

	return even
}

func main(){

	fmt.Println(fibonacci(50))
}
