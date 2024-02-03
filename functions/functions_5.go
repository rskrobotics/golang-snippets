package functions

import (
	// "fmt"
	"log"
	"math/rand"
)

func randomizedFunction() func() {
	if rand.Intn(2) == 0 {
		return func() { log.Println("YEAH") }
	} else {
		return func() { log.Println("OOOH YEAH") }
	}
}

func firstClassCitizens() {
	x := func() {
		// log.Println("Testing anonymous function")
	}

	func(s string) {
		log.Printf("Value that has been provided is %s", s)
	}("YIPPIE")

	x()

	for i := 0; i < 10; i++ {
		randomizedFunction()()
	}
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func operate(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	firstClassCitizens()
	x := operate(5, 9, add)
	log.Println(x)

	log.Printf("%T \t \n", add)
	log.Printf("%T \t \n", substract)
	log.Printf("%T \t \n", operate)

}
