package main

import (
	"errors"
	"fmt"
	"log"
)

//http://golang.prg/src/pkg/bufio/bufio.go
//http://golang.prg/src/pkg/io/io.go

//ErrNorgateMath this is idiom to start it with 'Err' since it used by the standard library.
var ErrNorgateMath = errors.New("Cannot square negative numbers ") //create custom error message and assign to variable

func main() {
	fmt.Printf("%T", ErrNorgateMath) //error string
	_, err := square(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func square(f float64) (float64, error) { //multi return values
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}
