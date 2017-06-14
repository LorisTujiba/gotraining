package strings

/*============================
	Exported and Not Exported
	==========================
	Exported 	= can be seen by another package
	Not Exp 	= can't be seen by another package

	Exported -> capitalized
	Not Exp -> lowercase
*/

var Name = "Joseph"
var status = "Single"

func GetStatus() string{
	return status
}
