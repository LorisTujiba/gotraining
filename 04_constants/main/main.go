package main

import "fmt"

/*==============================
const - Constant
============================
A simple value that cannot
be changed.
*/

const bebanHidup = "death & taxes"

/*===============================
iota & Define multiple constant
===============================
Increment, based on line.
reset on different
block.
*/

const (
	language = "GO"
	gender   = "Male"
	//A is Exported
	A = iota //2
	//B is Exported
	B = iota //3
	//C is Exported
	C       = iota //4
	address = "Kebon Jeruk"
	//D is Exported
	D = iota //6
)

const (
	//E is exported
	E = iota
	//F is exported
	F
)

func main() {

	const pi = 3.14

	//pi++
	//gender = "Female"

	fmt.Println(bebanHidup)
	fmt.Println(language)
	fmt.Println(gender)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

}
