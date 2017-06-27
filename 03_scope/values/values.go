package values

import "fmt"

/*============================
Exported and Non exported
==========================
Exported 	= can be seen by another package
Non Exp 	= can't be seen by another package

Exported -> capitalized
Non Exp -> lowercase

bukan public private disini
*/

//Name is Exported
var Name = "Joseph"
var status = "Single"

//GetStatus act as a getter for the unexported status
func GetStatus() string {
	return status
}

//PrintSeparation used to print separation
func PrintSeparation(input string, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(input)
	}
	fmt.Println()
}
