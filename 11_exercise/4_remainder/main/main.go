package main

import "fmt"

func main() {
	var smallNumber int
	var largerNumber int

	fmt.Print("Input number : ")
	fmt.Scan(&smallNumber)

	for largerNumber <= smallNumber {
		fmt.Print("Input larger number than before : ")
		fmt.Scan(&largerNumber)
	}

	fmt.Printf("Here is the remainder : %v\n", largerNumber%smallNumber)
}
