package main

import "fmt"

func main() {

	/*==============================
	  Array
	  ============================
	  Sequence of elements of a single
	  type. not dynamic, unlike
	  slice. Slice is dynamic.
	*/

	var numbers [10]int //array
	var scores []int    //slice

	numbers[4] = 4

	for i, v := range numbers {
		fmt.Println(i, " ", v)

	}

	scores = append(scores, 10, 11, 12, 31, 42)

	for i, v := range scores {
		fmt.Println(i, " ", v)
	}
}
