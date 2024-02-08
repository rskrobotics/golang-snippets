package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("GONNA EXIT!")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v, ok := <-e:
			if !ok {
				e = nil // Prevent further selects on this channel
			} else {
				fmt.Println(v, " from the even channel")
			}
		case v, ok := <-o:
			if !ok {
				o = nil // Prevent further selects on this channel
			} else {
				fmt.Println(v, " from the odd channel")
			}
		case v := <-q:
			fmt.Println(v, " from the quit channel")
			return
		}
		if e == nil && o == nil {
			break // Exit the loop if both even and odd channels are closed
		}
	}
}
