package main

import (
	"fmt"
	"sync"
	"github.com/LorisTujiba/gotraining/3_scope/values"
)

func main(){
	/*===========================================
	N-to-1
	===========================================
	Many function writing to one channel.
	*/

	c := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)//add 2 , so limit go routines into 2 only
	//the n
	go func (){
		for i:=0;i< 10;i++{
			c <- i
		}
		wg.Done()
	}()
	go func (){
		for i:=0;i< 10;i++{
			c <- i
		}
		wg.Done()
	}()

	go func(){
		wg.Wait()//waiting for the waitgroup is done executiong, then run this block of code
		close(c)
	}()


	//the one
	for n :=range c{//waiting to receive the baton, waits there till somebody sends on the channel
		fmt.Print(n)
	}

	fmt.Println()
	values.PrintSeparation("=",40)

}