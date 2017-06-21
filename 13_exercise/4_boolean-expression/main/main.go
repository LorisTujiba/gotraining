package main

import "fmt"

func main() {
	fmt.Println((false && true) || (true && false) || !(false && false))
}
