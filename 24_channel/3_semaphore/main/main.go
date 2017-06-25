package main

import "fmt"

func main(){
	/*===========================================
	Semaphore, like some kind of flag
	===========================================
	Plain variable that is changed depending
	on defined conditions. this variable
	later will be used as a condition
	to control access to some
	system resource.

	let's be purest, and only use channels. get rid of the wait group, mutex or atomicity
	*/

	c := make(chan int)
	done := make(chan bool)

	go func (){
		for i:=0;i< 10;i++{
			c <- i
		}
		done <- true //put true on the channel
	}()
	go func (){
		for i:=0;i< 10;i++{
			c <- i
		}
		done <- true //put true on the channel
	}()

	go func(){
		//throw value away
		<-done//waiting to receive the baton, waits there till somebody sends on the channel, if there's a value, ok this is finish
		//now this code below is waiting
		<-done//waiting to receive the baton, waits there till somebody sends on the channel, if there's a value, ok this is finish
		close(c)
	}()

	/*IF, we don't put the <- done into a go routine, what happened ?
	since this is an unbuffered channel, channel will accept one value. but only able to put the value when something
	is there to receive the value. They have to synchronize. Like the relay and the runner, they have to pass
	the baton hand to hand. So they have to be running at the same time, if they're not, we have a deadlock*/

	/*
		<-done  we're blocked right here
		<-done
		close(c)

	Alternative

	go func(){
		for i:=0;i<n;i++{ //n is the number of the routines
			<-done
		}
		close(c4)
	}()

	*/

	for n :=range c{//waiting to receive the baton, waits there till somebody sends on the channel
		fmt.Print(n)
	}
}