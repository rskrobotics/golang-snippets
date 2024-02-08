package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

func main() {
	///FANIN
	// even := make(chan int)
	// odd := make(chan int)
	// fanin := make(chan int)

	// go send(even, odd)

	// go receive(even, odd, fanin)

	// for v := range fanin {
	// 	fmt.Println(v)

	// }

	// fmt.Println("WE DONE BOYO")


	///FANOUT
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v:= range c2{
		fmt.Println(v)
	}
	
	fmt.Println("WE DONE AGAIN")

}

func populate(c chan int){
	for i:=0; i<100; i++{
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int){
	var wg sync.WaitGroup
	for v := range c1{
		wg.Add(1)
		go func(v2 int){
			c2<- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int{
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}


func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
