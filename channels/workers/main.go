package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	numWorkers := 10
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println("FROM C: ", v)
	}

}

func worker(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < 10; i++ {
		c <- rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
}
