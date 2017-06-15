package main

import "fmt"

/*==============================
	const - Constant
	============================
	A simple value that cannot
	be changed.
*/

const bebanHidup = "death & taxes"

/*--------------------------
	Define multiple constant
	------------------------
*/

/*==============================
	iota
	============================
	Increment, based on line.
	reset on different
	block.
*/

const (
	language = "GO"
	gender = "Male"
	A = iota //2
	B = iota //3
	C = iota //4
	address = "Kebon Jeruk"
	D = iota //6
)

const (
	E	= iota
	F
)

func main(){

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