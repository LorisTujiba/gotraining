package main

import "fmt"

func main() {

	a := 65

	fmt.Printf("Value \t\t\t\t\t: %v\n", a)
	fmt.Printf("Address \t\t\t\t: %v\n", &a)

	var pointOfA *int = &a

	fmt.Printf("Value of Pointer (memaddress) \t\t: %v\n", pointOfA)
	//Dereferencing
	fmt.Printf("Value of Pointer (pointed value)\t: %v\n", *pointOfA)
	fmt.Printf("Address of Pointer \t\t\t: %v\n", &pointOfA)

	fmt.Println("======================================================")

	//Everything in GO is pass by value!
	b := 5
	fmt.Printf("B address \t\t\t\t: %v\n", &b) //address
	reset(b)
	fmt.Printf("B value \t\t\t\t: %v\n", b) //still 5

	resetByMemoryAddress(&b)
	fmt.Printf("B value \t\t\t\t: %v\n", b) //goes 0

}

func reset(input int) {
	fmt.Printf("Reset's variable address \t\t: %v\n", &input) //differ from the main,
	// cuz its passed by value.
	// Not reference
	input = 0
}

func resetByMemoryAddress(input *int) {
	fmt.Printf("ResetByMemory's variable address \t: %v\n", &input) //check address
	*input = 0
}
