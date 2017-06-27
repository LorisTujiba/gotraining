package expression

import "fmt"

/*=====================
Function Expression
===================
Assign a function to
variable. Have func
inside a func.
*/

//MakeGreet is exported
func MakeGreet() func() string {
	return func() string {
		return fmt.Sprint("Returning a function that has string return type")
	}
}
