package main

import (
	"fmt"
	"strings"
)

func main() {

	/*=========================
	Switch
	=======================
	no fallthrough, no need
	to break each case
	*/

	var input int = 0
	for input < 1 || input > 7 {
		fmt.Print("Input day number : ")
		fmt.Scan(&input)
	}

	switch input {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Sunday")
	}

	/*=========================
	With fallthrough
	=======================
	*/

	var name string
	for len(strings.Trim(name, " ")) == 0 {
		fmt.Print("Input name : ")
		fmt.Scan(&name)
	}

	switch name {
	case "Even":
		fmt.Println("I know him")
		fallthrough //execute the next statement
	case "Agus", "Rio": //multiple condition
		fmt.Println("Both Agus and Rio are 14-0")
	case "Dianto", "Clinford":
		fmt.Println("Both Dianto and Clinford are my thesis team")
	default:
		fmt.Println("Dont know this guy")
	}

	switchWithType("string")
	switchWithType(450)
}

/*=========================
Switch on type
=======================
*/

func switchWithType(anonimus interface{}) {
	switch anonimus.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("This is an string")
	case float64:
		fmt.Println("This is an float")
	default:
		fmt.Println("Unknown")
	}
}
