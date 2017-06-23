package main

import "fmt"

func main() {

	/*====================================
	:= - Shorthand
	==================================
	Define variable without data types
	Initialization is a must
	*/

	name := "Loris"
	age := 22
	height := 164.2
	isaMale := true

	fmt.Printf("Hello!\nMy name is %v\nI'm %v years old\n", name, age)
	fmt.Printf("Height : %v\n", height)
	fmt.Printf("Male : %v\n", isaMale)

	/*===================
	Determine the types
	===================
	*/

	fmt.Printf("Name : %T\n", name)
	fmt.Printf("Age : %T\n", age)
	fmt.Printf("Height : %T\n", height)
	fmt.Printf("Gender : %T\n", isaMale)

	/*====================================
	var - variable
	==================================
	Define variable, automatically
	set the value to null if no
	initialization
	*/

	var brand string = "Mango"

	fmt.Println(brand)

}
