package values

/*============================
Exported and Non exported
==========================
Exported 	= can be seen by another package
Non Exp 	= can't be seen by another package

Exported -> capitalized
Non Exp -> lowercase
*/

var Name = "Joseph"
var status = "Single"

func GetStatus() string {
	return status
}
