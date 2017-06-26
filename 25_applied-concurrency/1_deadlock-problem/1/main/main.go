package main

import "fmt"

/*==============================================================
Case
==============================================================

func main(){
	c := make(chan int)
	c <- 1 #the case
	fmt.Println(<-c)
}

This results in a deadlock.
Can you determine why?
And what would you do to fix it?
*/

/*======
Answer
======
*/

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

/*--------------------------------------
Because, to be able to receive or pass a value,
the channel has to be ready and synchronize.
in the case above both of them not
synchronized. We have a runner
(c <-1), but there was nobody
to receive it. so we have a
code block. resulting
in deadlock.

To solve it, give it a go
command inside a anon
self-exe function.
--------------------------------------

*/
