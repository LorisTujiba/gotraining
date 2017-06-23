package main

import (
	"sort"
	"fmt"
)

/*===================================================
Package Sort
===================================================
Package sort provides primitives for sorting slices
and user defined collections (type). Jadi ttp
harus buat isi nya sendiri. Mirip adapter
di java android. Cuman sediain template
yang umum dan bisa di implement obj
lain.
 */

type subjectCoordinator []string

/*=======================
Sorting slice of string
=======================
sort exercise.

type Interface
A type, typically a collection, that satisfies sort.Interface can be sorted by
the routines in this package. The methods require that the elements of the
collection be enumerated by an integer index.

type Interface interface {
		// Len is the number of elements in the collection.
		Len() int
		// Less reports whether the element with
		// index i should sort before the element with index j.
		Less(i, j int) bool
		// Swap swaps the elements with indexes i and j.
		Swap(i, j int)
}

jadi bisa pakai fungsi itu untuk implement sort package
*/

func (input subjectCoordinator) Len() int{//Len dari sort package, jadi skrg subjectCoordinator implementing sort package
	return len(input)
}

func (input subjectCoordinator) Swap(i,j int){//dari sort package
	input[i],input[j] = input[j],input[i]
}

func (input subjectCoordinator) Less(i, j int) bool{//dari sort package
	return input[i] < input[j]
}

func main(){

	subcos := subjectCoordinator{"Loris","Agustian","William","Rio","Hengky","Hendra","Even","Goldwin","Christina"}
	/*--------------------------------------------------------------
		type StringSlice
			func (p StringSlice) Len() int
			func (p StringSlice) Less(i, j int) bool
			func (p StringSlice) Search(x string) int
			func (p StringSlice) Sort()
			func (p StringSlice) Swap(i, j int)

		cause StringSlice underlying type is string. If you have
		the string slice, we can access those methods. kalau
		string of slice ga bisa.
	--------------------------------------------------------------*/
	sort.Sort(subcos)//sort terima parameter tipe interface, type salah satunya
	fmt.Println(subcos)

	s := []string{"Zeno","John","Al","Jenny"}//ini slice of string, bukan string slice
	/*=====
	Cara 1
	======*/
	sort.StringSlice(s).Sort()//Convert s to string slice from slice of string
	//so this object is implementing sort package. Mana yang di implement? yang String Slice
	//Jadi bisa pakai fungsi-fungsi StringSlice one of which is Sort. Kalau ga diubah kan jadi ga punya len, swap
	//Jadi ga bisa sort
	//sort.Sort(s) cannot use s as type interface
	fmt.Println(s)
	fmt.Printf("%T",s)//bentuknya slice of string

	/*=====
	Cara 2
	======*/
	st := []string{"Zeno","John","Al","Jenny"}
	sort.Sort(sort.StringSlice(st))//sort.Sort itu dari package, terus argumen nya kita ubah si s(slice of String) jadi String Slice.
	fmt.Println(st)

	//buktinya
	st2 := sort.StringSlice(st)
	fmt.Printf("%T",st2)//balikin nya StringSlice. knp di ubah ke StringSlice?
	//biar bisa implement package sort. liat di documentasi atas tadi.

	/*=====
	Cara 3
	======
	----------------------------------------------------------
	func Strings(a []string)// yang ini terima slice of string
	------------------------------------------------- ---------
	*/
	str := []string{"Zeno","John","Al","Jenny"}
	sort.Strings(str)
	fmt.Println(str)

	/*=======================
	Reverse
	=======================
	Reverse exercise
	*/
	//reverse subco (type)
	sort.Sort(sort.Reverse(subcos))//sort need an interface, so we give the reverse
	//krn reversenya balikin interface
	fmt.Println(subcos)

	//reverse slice of string
	//sort.Sort() ini ga balikin apa apa, tapi butuh interface sbg parameter
	//sort.Reverse balikin interface jadi bisa kalau mau ditaruh, tapi butuh data interface

	//sort.Sort(sort.Reverse(s)) sayangnya s itu slice of string, jadi ubah dulu
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	//kenapa ga bisa langsung reverse ? soalnya balikin nya addressnya
	//Reverese overrides less method and returns the address of the modified interface
	i := sort.Reverse(sort.StringSlice(st))
	fmt.Println(i)
	fmt.Printf("%T",i)//balikin nya pointer sort.reverse





//n:= []int{7,4,8,2,9,19,12,32,3}
}