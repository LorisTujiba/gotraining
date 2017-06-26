package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*==================
	Conversion
	==================
	changing data type
	*/
	var x int = 21
	var y float64 = 60.2

	fmt.Println(x * int(y))     //float to int conversion (narrowing)
	fmt.Println(float64(x) * y) //int to float conversion (widening)

	fmt.Println(string([]byte{'H', 'i', ' ', 'h', 'e', 'l', 'l', 'o'})) //rune to string conversion
	fmt.Println([]byte("Yo"))                                           //string to rune conversion

	z, _ := strconv.Atoi("65") //ascii to int
	fmt.Println(z)
	fmt.Println("This is ", strconv.Itoa(65)) //int to ascii

	/*=====================
	Conversion with Parse
	=====================
	*/

	b, _ := strconv.ParseBool("true")
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	f, _ := strconv.ParseFloat("42.34", 64)

	fmt.Println(b, i, u, f)

	/*=====================
	Conversion with Parse
	=====================
	*/

	fb := strconv.FormatBool(true)
	fi := strconv.FormatInt(-23, 10)
	ff := strconv.FormatFloat(234.2, 'g', -1, 64)
	fu := strconv.FormatUint(23, 10)

	fmt.Println(fb, fi, ff, fu)

	/*===================
	Assertion
	===================
	For type interfaces
	int(v)<-conversion,dikiri
	v.(int)<-assertion,dikanan
	*/
	//nameTest := "Clinford"
	//str,ok_ := nameTest.(string) kalau bukan interface ga bisa

	var name interface{} = "Test"
	//var name interface{} = 32

	string, ok := name.(string) //hey, is name a string?

	if ok {
		fmt.Printf("%T\n", string)
	} else {
		fmt.Printf("Not a string, it's %T\n", string)
	}

	var age interface{} = 10

	//fmt.Println(age * 30)//invalid operation, mismatch type, harus precise. jadi di assert dulu
	fmt.Println(age.(int) * 2) //di assert baru jalan

}
