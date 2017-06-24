package main

import (
	"fmt"
	"github.com/LorisTujiba/gotraining/3_scope/values"
)

/*=========================================
type [name] struct - Struct #1
=========================================
Create your own type and assign value.
This type can have fields. In Go,
we don'tcreate a classes. we
create a type. we don't
instantiate, we create
value of a type. GO
is static type.

type [name] underlying type #2
*/

type human struct { // declaring the type
	name   string //this is the list of the fields
	age    int
	height float64
	weight float64
}

//Create Method
func (input human) greeting() string {
	return fmt.Sprint("My name is ", input.name, " Im ", input.age, " years old")
}

func (input human) eat() string {
	return fmt.Sprint("Im eating")
}

/*=================
Composition
=================
Type inside a
type.
*/

type staff struct {
	human
	age    int // field override
	salary float64
}

//Method Overriding
func (input staff) greeting() string {
	return fmt.Sprint("Hi, Im a Staff ", input.age, " my salary : ", input.salary)
}

type codeName int

/*----------------------------------------
Printing - effective go
If you want to control the default format
for a custom type, all that's required is
to define a method with the signature
String() string on the type. For our
simple type T, that might look like
this.

func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
------------------------------------------*/

func (s human) String() string {
	return fmt.Sprintf("My name is %s and im %d years old", s.name, s.age)
}

func main() {

	var myCode codeName

	myCode = 12

	var Agustian = human{ //struct literal
		"Agustian Rinaldy",
		22,
		120.2,
		165.3,
	}

	fmt.Printf("%v\n", Agustian) //kena di fungsi string
	fmt.Println(Agustian)        //sama kena juga
	fmt.Println("Hi, my name is ", Agustian.name)
	fmt.Println("I'm ", Agustian.age, " years old!")
	fmt.Println("This is the on using underlying type : ", myCode)
	values.PrintSeparation("=", 50)
	fmt.Println(Agustian.eat())

	Even := staff{
		human: human{
			name:   "Even Yosua",
			age:    22,
			height: 160,
			weight: 60,
		},
		age:    23,
		salary: 400000,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(Even.eat())               //Even can use method eat
	fmt.Println(Even.greeting())          //Overriding
	fmt.Println(Even.age, Even.human.age) //field override

	//Struct pointer

	var Goldwin = &staff{
		human: human{
			name:   "Goldwin Japar",
			age:    22,
			height: 166,
			weight: 57,
		},
		age:    22,
		salary: 400000,
	}

	fmt.Println(*Goldwin)
	fmt.Println(Goldwin)
	fmt.Println(Goldwin.age)
	fmt.Println(Goldwin.name)
}
