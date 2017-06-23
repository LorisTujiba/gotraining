package main

import "fmt"

func main() {

	/*=========================
	If
	=======================
	Simply a selection to
	control the flow
	of the program
	================================
	Boolean Expression
	==============================
	&& both condition must be true
	|| one of them must be true
	! negate
	*/

	if a := 60; a > 0 && a <= 60 { //if with declaration inside
		fmt.Println("a is between 0 and 60")
	} else if a > 60 && a < 100 {
		fmt.Println("a is between 61 and 100")
	} else {
		fmt.Println("other")
	}
}
