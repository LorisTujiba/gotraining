package main

import "fmt"

/*==============================================================
Case
==============================================================
func main(){
	c := make(chan int)

	go func(){
		for i:=0;i<10;i++{
			c <- i
		}
		close(c)
	}()

	fmt.Println(<-c)
}

Why does this only print zero?
And what can you do to get it to print all 0 - 9 numbers ?
*/

/*======
Answer
======
*/
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c { //keep waiting for a value till it closed
		fmt.Println(n)
	}

}

/*
Because the receiver is only one. so we only receive
the first value. Which is 0. after that the main
exit.
to solve, use range to loop all it's content
and print it
*/
