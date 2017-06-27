package passbyvalue

import "fmt"

/*===================================
Pass by Value
==================================
Go always pass by value. Therefore
the value does not carried to
the main. because they have
different address. Vice
versa if we passed the
address by value.

If it's not a reference type (i.e int, string)
and you want to change the value, you need to
pass the address. But if it's a reference
type, you don't need to pass the address
because its already referenced.
*/

//Aging is exported
func Aging(input int) {
	fmt.Println("Calling from aging function : ", input)
	input++
	fmt.Println("Change the age: ", input)
}

//AgingReference calling with reference type
func AgingReference(input *int) {
	fmt.Println("Calling from aging reference function : ", *input)
	*input++
	fmt.Println("Change the age reference : ", *input)
}

//ReferencedType , slice of string is a referenced type
func ReferencedType(input []string) {
	input[0] = "Loris"
	input[1] = "Tujiba"
}
