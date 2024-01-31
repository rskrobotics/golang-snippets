package main

import (
	"fmt"
)

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

func (d make) String() string {
	return [...]string{"Toyota", "Skoda", "VW"}[d]
}

// func (e engine) String() string {
// 	if e.electric{
// 	return "Electric"
// 	}
// 	return "Combustion"
// }

// func (v vehicle) String() string {
//     return fmt.Sprintf("Make: %s, Engine: %s", v.make, v.engine)
// }

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
	second_excercise()
}