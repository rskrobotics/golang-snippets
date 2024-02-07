package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS \t\t", runtime.GOOS)
	fmt.Println("ARCH \t\t", runtime.GOARCH)
	fmt.Println("CPU's \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU's \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	i := 10
	for i > 0 {
		fmt.Println("FOO: ", i)
		i--
	}
	wg.Done()
}

func bar() {
	i := 10
	for i > 0 {
		fmt.Println("BAR: ", i)
		i--
	}

}
