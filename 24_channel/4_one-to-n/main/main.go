package main

import (
	"fmt"
)

func main(){
	/*===========================================
	1-to-N
	===========================================
	Many function pulling from the same channel
	we have 2 n. 2 func pulling.
	*/

	c := make(chan int)
	done := make(chan bool)

	go func (){
		for i:=0;i< 10000;i++{
			c <- i
		}
		close(c)
	}()

	//n 1
	go func(){
		for n :=range c{//waiting to receive the baton, waits there till somebody sends on the channel
			fmt.Println(n)
		}
		done <- true
	}()

	//n2
	go func(){
		for n :=range c{//waiting to receive the baton, waits there till somebody sends on the channel
			fmt.Println(n)
		}
		done <- true
	}()

	//holding main from exit, wait till the n is done
	<-done
	<-done
}
