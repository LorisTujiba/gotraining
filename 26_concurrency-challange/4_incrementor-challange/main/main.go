package main

import (
	"fmt"
	"strconv"
)

/*
import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}


CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/

func main() {
	c := incrementor(2)

	for n := range c {
		fmt.Println(n)
	}
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				out <- fmt.Sprint("Process : ", strconv.Itoa(i), " Printing : ", j)
			}
			done <- true
		}(i)
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()

	return out
}
