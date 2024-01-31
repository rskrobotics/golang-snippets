package main 

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
	first string
}

func main(){
	fmt.Println("Yeah!!")

	// p1 := person{
	// 	first: "James",
	// 	last: "Bond",
	// 	age: 32,
	// }

	sa1 := secretAgent{
		person: person{
			first: "REKT",
			last: "Shrekt",
			age: 47,
		},
		first: "GANDALF THE BLUE",
		ltk: true,
	}



	// fmt.Println(p1)
	// fmt.Println(sa1)
	// fmt.Printf("%T \t %v \n", p1, p1.last)
	// fmt.Printf("%T \t %v \n", sa1, sa1.ltk)
	fmt.Printf("%#v \n", sa1.first)
	fmt.Printf("%#v \n", sa1.person.first)
}