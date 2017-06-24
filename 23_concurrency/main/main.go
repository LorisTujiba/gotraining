package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

/*===================================
go - Concurrency & Parallelism
===================================
Is the composition of independently
executing process.
command : go
While, parallelism is like
doing many things at the
same time.
*/

var wait sync.WaitGroup //batasin 2 aja yang jalan
var mutex sync.Mutex
var counter int64

func main(){

	//go foo()
	//go john() can't run, because there can be only 2 process simultaneously. Currently there are 3 plus the main
	//therefore, we use the waitgroup

	wait.Add(2)
	go foo()
	go john()
	wait.Wait()

	/*=====================================
	Race Condition
	=====================================
	Condition where two process accessing
	the same variables, resulting in
	overwriting the variable. so,
	we can use mutex to solve
	this condition.
	*/
	wait.Add(2)
	go incrementor("Foo :")
	go incrementor("Bar :")
	wait.Wait()
	fmt.Println("Final Counter : ",counter)

	/*=====================================
	Atomicity
	=====================================
	Control access for variable, so only
	one can access a variable.
	*/
	counter = 0
	wait.Add(2)
	go incrementorAtomic("Foo :")
	go incrementorAtomic("Bar :")
	wait.Wait()
	fmt.Println("Final Counter : ",counter)
}

func incrementor(s string){
	for i:= 0;i<20;i++{
		time.Sleep(time.Duration(rand.Intn(20))* time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s,i,"Counter",counter)
		mutex.Unlock()
	}
	wait.Done()
}


func incrementorAtomic(s string){
	for i:= 0;i<20;i++{
		time.Sleep(time.Duration(rand.Intn(20))* time.Millisecond)
		atomic.AddInt64(&counter,1)
		fmt.Println(s,i,"Counter",counter)
	}
	wait.Done()
}


func foo(){
	for i:=0;i<100;i++{
		fmt.Println("foo : ",i)
		time.Sleep(3000)
		//time.Sleep(time.Duration(3*time.Millisecond)
	}
	wait.Done()
}

func john(){
	for i:=0;i<100;i++{
		fmt.Println("john : ",i)
		time.Sleep(3000)
	}
	wait.Done()
}