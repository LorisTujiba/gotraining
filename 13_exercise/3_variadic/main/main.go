package main

import "fmt"

func findMax(inputs ...int) int {
	var max = 0
	for _,input := range inputs{
		if input > max{
			max = input
		}
	}
	return max
}

func main(){
	fmt.Println(findMax(10,2,4,6,11,19,16,23,4))
}
