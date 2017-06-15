package main

import "fmt"

func main(){

	/*===============================
		Memory Address
		=============================
		Use mem address to save value.
	*/
	var a int
	fmt.Print("Input some value : ")
	fmt.Scan(&a)
	fmt.Printf("Value : %v - Address : %v\n",a,&a)

}