package main

import (
	"encoding/json"
	"fmt"
)

type ketchup struct {
	Sugar_amt float32
	Spicy     bool
}

type MenuItem struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Price       float64  `json:"cena"`
    Ingredients []string `json:"ingredients"`
    Vegan       bool     `json:"vegan"`
}

func main() {
	incoming_data := `{
		"id": "item123",
		"name": "Veggie Pizza",
		"cena": 12.99,
		"ingredients": ["Tomato Sauce", "Mozzarella", "Mushrooms", "Bell Peppers", "Onions"],
		"vegan": false
	}`
	mi := &MenuItem{}
	err := json.Unmarshal([]byte(incoming_data), mi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v",*mi)

	// k1 := ketchup{
	// 	Sugar_amt: 31.2,
	// 	Spicy:     false,
	// }
	// k2 := ketchup{
	// 	Sugar_amt: 4.7,
	// 	Spicy:     true,
	// }

	// ketchup_tray := []ketchup{k1, k2}

	// fmt.Println(ketchup_tray)

	// bs, err := json.Marshal(ketchup_tray)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bs), err)

	

}
