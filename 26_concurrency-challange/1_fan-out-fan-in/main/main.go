package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

//Question
//1. Is this Fan Out ?
//yes, fan out means we distribute same channel into several function
//this can be seen in line first line of the main, which we declared a channel
//then we distribute it using the workerProcess function three times

//2. Is this Fan in ?
//many channels merged into one channel. we don't have many channels here
//so this not fan in

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * 1000)

}

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is :%s\n", thisID, input)
	}
}
