package main

import (
	"fmt"
	"time"
)

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func factorial(n int ) int {
	if n == 0{ return 1}
	return n * factorial(n-1)
}

func busyWork(){
	for i:=0; i<200000; i++{
		fmt.Println(i)
	}
}

func profiler(f func()) {
	start := time.Now()
	f()
	fmt.Println("Time since start: ", time.Since(start))
}

func main() {
	profiler(busyWork)
	// f := igo rncrementor()
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	// fmt.Println(factorial(8))
	
}


