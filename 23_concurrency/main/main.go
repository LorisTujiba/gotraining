package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*===================================
go - Concurrency & Parallelism
===================================
Is the composition of independently
executing process. Dealing with
lots of things at once.
command : go
While, parallelism is like
doing many things at the
same time. Doing lots
of things at once.

Concurrency is not Parallelism.
*/

var wait sync.WaitGroup //batasin 2 aja yang jalan
var mutex sync.Mutex
var counter int64

func main() {

	//go foo()
	//go john()Currently there are 3 routines, main, foo bar. 3 of them run simultaneously. But the main ended first.
	//so, nothing happened. therefore, we use the waitgroup to hold the func main

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
	we can use mutex (mutually
	exclusive) to solve
	this condition.

	run with --race to check if there's
	a race condition
	*/
	wait.Add(3)
	go incrementor("Foo :")
	go incrementor("Bar :")
	go incrementor("john :")
	wait.Wait()
	fmt.Println("Final Counter : ", counter)

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
	fmt.Println("Final Counter : ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter", counter)
		mutex.Unlock()
	}
	wait.Done()
}

func incrementorAtomic(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter", counter)
	}
	wait.Done()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo : ", i)
		time.Sleep(3000)
		//time.Sleep(time.Duration(3*time.Millisecond)
	}
	wait.Done()
}

func john() {
	for i := 0; i < 100; i++ {
		fmt.Println("john : ", i)
		time.Sleep(3000)
	}
	wait.Done()
}
