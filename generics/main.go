package main

import "fmt"

func swapper[T, V any](a T, b V)(V, T){
	return b, a
}

func main(){
	a := 3
	b := 6.2
	c, d := swapper(a,b)
	fmt.Println(c,d)

}