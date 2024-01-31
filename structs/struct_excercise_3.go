package main

import (
	"fmt"
)



func main(){
	p1 := struct{
		name string
		friends map[string]int
		favDrinks []string
	}{
		name: "KK",
		friends: map[string]int{
			"Soniak": 30,
			"Krystian": 29,
			"Waldemar": 33,
		},
		favDrinks: []string{
			"MILK", "COLA",
		},
	}


	fmt.Println(p1)
}