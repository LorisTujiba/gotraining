package main

import "fmt"

/*==============================
Slice
============================
A referenced type. Slice is
dynamic. It's indexable
and have a length.
*/

func main(){
	theSlice := []int{12,13,14,15,16,17,18,19}//shorthand slice
	fmt.Println(theSlice)//print all
	fmt.Println(theSlice[2:5])//print from index 2 to 4, 5 not included
	fmt.Println(theSlice[2])//print at index 2
	fmt.Println("get this A's ASCII"[9])//get string's char at index of 9. which is A. and ASCII of A is 65

	anotherSlice := make([]int,0,3)//create a slice that has 3 capacity and 0 len

	for i:=0;i<20;i++{
		anotherSlice = append(anotherSlice, i)
		fmt.Println("capacity : ",cap(anotherSlice),"Len : ",len(anotherSlice))
	}//if the len exceeds the capacity, go will twice the capacity

	//fmt.Println(anotherSlice[51])//index out of range, only use append if you want to extend the capacity
	theSlice = append(theSlice,anotherSlice...)//appending 2 slice
	fmt.Println(theSlice)

	//Slice of slice of string

	records := make([][]string,0)

	studentOne := make([]string,2)
	studentOne[0]="Joseph"
	studentOne[1]="A+"

	records = append(records,studentOne)

	studentTwo := make([]string,2)
	studentTwo[0] = "Agustian"
	studentTwo[1] = "C"

	records = append(records,studentTwo)

	fmt.Println(records)

}
