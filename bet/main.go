package main

import (
	"fmt"
	"github.com/rskrobotics/golang-snippets/bet/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}


// Gen generates! :D
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 0, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	return [][]int{a, b, c, d}
}
