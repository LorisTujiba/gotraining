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

var Name = "Joseph"
var status = "Single"

func GetStatus() string {
	return status
}

func PrintSeparation(input string, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(input)
	}
	fmt.Println()
}
