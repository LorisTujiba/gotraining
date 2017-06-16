package main

import (
	"fmt"
	"github.com/LorisTujiba/gotraining/3_scope/values"
)

/*==============================
Package level scope
============================
Can be accessed by the whole
corresponding file
*/
var number = 12

func main() {

	fmt.Println(number)
	fmt.Println(calculate(7))
	fmt.Println("I can access this because it is exported : " + values.Name)
	fmt.Println("I can access this because the getter is exported : " + values.GetStatus())
}

/*==============================
Block level scope
============================
Can be accessed by inner
parentheses. cant be
accessed by outer
parentheses.
*/
func calculate(input int) int {
	return input + number
}
