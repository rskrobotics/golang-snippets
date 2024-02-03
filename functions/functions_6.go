package main

import (
	"testing"
	"fmt"
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

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(factorial(8))

}

func TestFactorial(t *testing.T){
	result := factorial(5)
	if result != 120{
		t.Errorf("Test incorrect, got: %d, want: %d", result, 120)
	}
}
