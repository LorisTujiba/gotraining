package main

import "fmt"

/*======================================
Interface
======================================
Interface is an type that only declare
behavior. This behavior never
implemented directly. but
instead by user-defined
type.
*/


type square struct{
	lebar float64
}

type triangle struct{
	alas float64
	tinggi float64
}

//Interface, support polymorphism-ish

type shape interface{
	area() float64 // every struct that have this function are a shape
}

func (s square) area() float64{// square implementing the shape
	return s.lebar * s.lebar
}

func (t triangle) area() float64{// triangle implementing the shape
	return (t.alas * t.tinggi)/2
}

func info( input shape){
	fmt.Println(input)
	fmt.Println(input.area())
}

func main(){

	kotak := square{50}
	segitiga := triangle{10,5}

	info(kotak)
	info(segitiga)

}