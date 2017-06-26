package main

import (
	"fmt"
	"github.com/LorisTujiba/gotraining/3_scope/values"
	"sync"
	"time"
)

/*===========================================
<- Channels
===========================================
Channels allow us to communicating between
go routines. Do not communicate by
sharing memory,instead, share
memory by communicating.
*/

func main() {

	c := make(chan int) //create channel that has int type, unbuffered channel
	//so far,make used to create maps, channels and slices only
	//these three types must be initialized before use

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // pass the counter to the channel //1
		}
	}() //anon self executed function

	go func() {
		for {
			fmt.Println("Channel value :", <-c) // channel : "give me whatever values there!" //2
		}
	}()

	time.Sleep(time.Second)

	//how does channel work ?
	//the code on the #1 stops until the #2 ready to take the value
	//and after that, if #2 ready to receive,#1 waiting till its ready to pass
	//think of it's like relay race, one waiting to get the baton, and the other one passing the baton
	//TLDR, passing data back and forth with some blocking nature

	/*===========================================
	Range
	===========================================
	Channels allow us to communicating between
	go routines. Do not communicate by
	sharing memory,instead, share
	memory by communicating.
	*/

	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}
		close(c2) //when we're done putting values to the channel, close the channel
		//when it's close, we cant put a value inside channel, but the channel
		//still receiving the value
	}()

	for n := range c2 {
		fmt.Print(n)
	}

	fmt.Println()
	values.PrintSeparation("=", 40)
}
