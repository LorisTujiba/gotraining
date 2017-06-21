package main

import "fmt"

func even(input int) (int, bool) {
	return input / 2, input%2 == 0
}

func main() {
	fmt.Println(even(1))
}
