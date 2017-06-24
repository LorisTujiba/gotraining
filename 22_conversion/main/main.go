package main

import (
	"fmt"
	"strconv"
)

func main(){
	var x int = 21
	var y float64 = 60.2

	fmt.Println(x * int(y))//float to int conversion (narrowing)
	fmt.Println(float64(x) * y)//int to float conversion (widening)

	fmt.Println(string([]byte{'H','i',' ','h','e','l','l','o'}))//rune to string conversion
	fmt.Println([]byte("Yo"))//string to rune conversion

	z,_ := strconv.Atoi("65")//ascii to int
	fmt.Println(z)
	fmt.Println("This is ",strconv.Itoa(65))//int to ascii

	/*==============
	Assertion
	==============
	For type
	*/
	var name interface{} = "Test"
	//var name interface{} = 32

	string, ok := name.(string)//hey, is name a string?

	if ok{
		fmt.Printf("%T\n",string)
	}else{
		fmt.Printf("Not a string, it's %T\n",string)
	}

	var age interface{} = 10

	//fmt.Println(age * 30)
	fmt.Println(age.(int) * 2)

}