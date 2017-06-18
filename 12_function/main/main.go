package main

import (
	"fmt"
	"github.com/LorisTujiba/gotraining/12_function/expression"
	"github.com/LorisTujiba/gotraining/12_function/variadic"
	"github.com/LorisTujiba/gotraining/12_function/callback"
)

func main() {
	/*===================
	Function
	=================
	Start with func keyword. Allow us to put
	code into container that can be used
	over and over again.
	---------
	Arguments
	---------
	value that passed to parameter
	---------
	Parameter
	---------
	required declaration for a function
	------
	Return
	------
	returning a value when the
	function is called
	*/

	greeting("Loris", "Tujiba") //arguments
	fmt.Println(getName("Putra"))
	fmt.Println(getNames("Doe", "John"))

	fmt.Println(variadic.Calculate(10, 20, 30, 40.5, 50.1))
	var datas = []float64{100.50, 21.56, 48.21, 913.2}
	fmt.Println(variadic.Calculate(datas...)) //passing variadic input

	/*------------------------
	Function Expression
	----------------------
	*/
	myFunc := expression.MakeGreet()
	fmt.Println(myFunc())

	myAnotherFunc := func() string {
		return "Hi"
	}
	myAnotherFunc()

	/*------------------------
	Function Callback
	----------------------
	*/
	total := callback.IsAMale([]string{"Male","Male","Female","Male"},func (input string)bool{
		return input == "Male"
	})
	fmt.Println(total)
}

func greeting(firstName, lastName string) {
	fmt.Println("Hello! My name is", firstName, lastName)
}

//return with var specifies
func getName(firstName string) (s string) {
	s = firstName
	return
}

//return a multiple values
func getNames(firstName, lastName string) (string, string) {
	return fmt.Sprint(firstName, lastName), fmt.Sprint(lastName, firstName)
}
