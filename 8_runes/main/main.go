package main

import "fmt"

func main() {
	/*======================
	Rune
	====================
	A character (numeric). Alias to int32 type, which
	is a numeric type to store 4 bytes. Because
	UTF-8 is a 4 byte coding scheme. Strings
	is a collection of runes.
	*/

	a := 'a'                    // returns 97
	fmt.Printf("%T\n", a)       // returns int32
	fmt.Printf("%T\n", rune(a)) // returns int32

	/*======================
	Conversion
	====================
	No "casting" in GO
	*/
	//convert string to slice of bytes
	fmt.Println([]byte("Hello")) //each char use 1 byte
	// 72 101 108 108 111
	// the rune, ascii
	fmt.Println([]byte("我愛你")) //each char use 3 byte
}
