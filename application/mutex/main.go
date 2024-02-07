package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

func main() {
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
			fmt.Println("Goroutines \t", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	fmt.Println("Counter \t", counter)

}
