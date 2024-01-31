package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	flavours []string
}

type engine struct{
	electric bool
}

type make int

type vehicle struct{
	engine
	make 
}

const (
	Toyota make = iota
	Skoda
	VW
)

// func (d make) String() string {
// 	return [...]string{"Toyota", "Skoda", "VW"}[d]
// }

// func (e engine) String() string {
// 	if e.electric{
// 	return "Electric"
// 	}
// 	return "Combustion"
// }

// func (v vehicle) String() string {
//     return fmt.Sprintf("Make: %s, Engine: %s", v.make, v.engine)
// }



func first_excercise(){
	p1 := person{
		first: "Sonia",
		last: "Śmiałek",
		flavours: []string{"Caramel","Toffee"},
	}
	p2 := person{
		first: "KK",
		last: "CRMO",
		flavours: []string{"VANILLA","COFFEE"},
	}

	// fmt.Printf("%T \t %v \n", p1, p1)
	// fmt.Printf("%T \t %v \n", p2, p2)
	// fmt.Println(p1.flavours)
	// fmt.Println(p2.flavours)

	// for _, v := range p1.flavours {
	// 	fmt.Println(p1.first, "favourite is ", v)
	// }

	m := map[string]person{}
	m[p1.last] = p1
	m[p2.last] = p2

	for _, v := range m{
		for _, f := range v.flavours{
			fmt.Println(v.first, "'s favourite is ", f)
		}
	}

}

func second_excercise(){

	v1 := vehicle{
		engine: engine{true},
		make: Toyota,
	}
	v2 := vehicle{
		engine: engine{false},
		make: Skoda,
	}

	fmt.Println(v1)
	fmt.Println(v2)
	
	fmt.Printf("%T \t %v \n", v1, v1)
	fmt.Printf("%T \t %v \n", v2, v2)



}

func main(){
	// first_excercise()
	second_excercise()
}