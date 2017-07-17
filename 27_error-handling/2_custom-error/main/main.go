package main

import (
	"errors"
	"log"
)

func main() {
	_, err := square(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func square(f float64) (float64, error) { //multi return values
	if f < 0 {
		return 0, errors.New("Cannot square negative numbers")
	}
	return 42, nil
}
