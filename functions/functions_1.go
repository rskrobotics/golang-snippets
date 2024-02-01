package main	

// func (r reciever) identifier (p parameter/s) (return(s)){ ... }

import (
	"fmt"
	"math/big"
)

func main()  {
	// foo()
	s := []int{1,2,3,7,9,67,455,23232,656653,24324}
	fmt.Println(product(s...))
}

func foo(){
	fmt.Println("BAR")
}

// Variadic parameters + defer

func product(input ...int) *big.Int {
	defer fmt.Println("EOF")
	res := big.NewInt(1)
	for _, v := range input{
		res.Mul(res, big.NewInt(int64(v)))
	}
	return res
}
