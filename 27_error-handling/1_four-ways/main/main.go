package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//go doesn't have exception, but instead we got multiple return value that can be used

	//error values in go aren't special, they are just values like any other, and so you have the
	// entire language at your disposal - Rob Pike
	_, err := os.Open("doesn'tExist.txt")
	if err != nil {
		fmt.Println("ouch, ", err)
		log.Println("ugh, ", err)
		log.Fatal("jeh, ", err)
		panic(err)
	}

}
