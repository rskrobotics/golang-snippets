package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			wg.Done()
			fmt.Println("Goroutines \t", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	fmt.Println("Counter \t", counter)

}
