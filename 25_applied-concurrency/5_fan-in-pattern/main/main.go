package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*========================================
Fan out/Fan in
========================================
A design pattern. Distributing the input data across some functions, this is called fanning
out. If we have channel inputs, we're going to fan it out. It means there are multiple
functions reading from that channel until it's closed. Fan In, means multiple
channels different values and write to the same channel.
*/

func main() {
	c := fanIn(boring("Agus"), boring("Andry"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring: I'm leaving.")
}

//this is not fan out because no channel coming in to this function
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string { //only receive channel that has the type of strin
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 //take the value of input1 and store it to c, waiting to receive value
		}
	}()
	go func() {
		for { //infinite loop
			c <- <-input2 //take the value of input2 and store it to c
		}
	}()
	return c
}
