package main

import "fmt"

func even(input int) func() (int, bool) {
	return func() (int, bool) {
		return input / 2, input%2 == 0
	}
}

func main() {
	result := even(2)
	fmt.Println(result())

	anotherResult := func(input int) (int, bool) {
		return input / 2, input%2 == 0
	}
	fmt.Println(anotherResult(6))

}
