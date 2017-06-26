package main

import (
	"fmt"
	"github.com/LorisTujiba/gotraining/20_interface/emptyinterface"
	"github.com/LorisTujiba/gotraining/3_scope/values"
)

/*======================================
Interface
======================================
Interface is an type that only declare
behavior. This behavior never
implemented directly. but
instead by user-defined
type.
*/

type square struct {
	lebar float64
}

type triangle struct {
	alas   float64
	tinggi float64
}

//Interface, support polymorphism-ish

type shape interface {
	area() float64 // every struct that have this function are implementing shape
	areaWithPointerReceiver() float64
}

func (s square) area() float64 { // square implementing the shape
	return s.lebar * s.lebar
}

func (t triangle) area() float64 { // triangle implementing the shape
	return (t.alas * t.tinggi) / 2
}

func (s *square) areaWithPointerReceiver() float64 {
	return s.lebar * s.lebar
}

func (t *triangle) areaWithPointerReceiver() float64 {
	return (t.alas * t.tinggi) / 2
}

func info(input shape) {
	fmt.Println(input)
	fmt.Println(input.area())
}

func infoWithPointerReceiver(input shape) {

	fmt.Println(input)
	fmt.Println(input.areaWithPointerReceiver())
}

func main() {

	kotak := square{50}
	segitiga := triangle{10, 5}

	info(&kotak)
	info(&segitiga)

	/*==============================
	Empty Interface
	=============================
	may hold values of any type
	usually handles values of
	unknown type.
	*/

	Labrador := emptyinterface.Dog{emptyinterface.Animal{"Labrador Retriever"}, 32}
	Pony := emptyinterface.Horse{emptyinterface.Animal{"Pony"}, 100}
	Shark := emptyinterface.Fish{emptyinterface.Animal{"Shark"}, true}

	animals := []emptyinterface.Animals{Labrador, Pony, Shark} // those types, can be stored inside animals
	//because everything implements animal
	animals2 := []interface{}{Labrador, Pony, Shark} //cara 2

	fmt.Println(animals)
	fmt.Println(animals2)

	/*==============================
	Receivers
	=============================
	Receivers       values
	(type Type)     T and *T works
	(type *Type)	*T only this work

	values			Receivers
	T				(t T)
	*T				(t T) and (t *T)
	*/
	values.PrintSeparation("=", 50)
	info(&kotak)
	values.PrintSeparation("=", 10)
	info(&kotak)
	values.PrintSeparation("=", 10)
	//infoWithPointerReceiver(kotak)// will not work
	infoWithPointerReceiver(&kotak) //this will work
}
