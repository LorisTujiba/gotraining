package main

import "fmt"

func foo(inputs ...int) {
	for _, input := range inputs {
		fmt.Println(input)
	}
}

func main() {
	//create a function that can be used with these way
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
