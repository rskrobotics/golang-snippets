package main

import (
	"encoding/json"
	"fmt"
	"sort"
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

type byVegan []MenuItem

func (b byVegan) Len() int      { return len(b) }
func (b byVegan) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byVegan) Less(i, j int) bool {
	if b[i].Vegan && !b[j].Vegan {
		return true
	}
	return false
}

func main() {
	incomingData := `[
        {
            "id": "item123",
            "name": "Veggie Pizza",
            "cena": 12.99,
            "ingredients": ["Tomato Sauce", "Mozzarella", "Mushrooms", "Bell Peppers", "Onions"],
            "vegan": false
        },
        {
            "id": "item124",
            "name": "Spicy Vegan Pizza",
            "cena": 14.99,
            "ingredients": ["Tomato Sauce", "Vegan Cheese", "Jalapenos", "Red Onions", "Vegan Sausage"],
            "vegan": true
        },
        {
            "id": "item125",
            "name": "Margherita Pizza",
            "cena": 11.99,
            "ingredients": ["Tomato Sauce", "Mozzarella", "Fresh Basil"],
            "vegan": false
        },
        {
            "id": "item126",
            "name": "BBQ Chicken Pizza",
            "cena": 15.99,
            "ingredients": ["BBQ Sauce", "Chicken", "Red Onions", "Cilantro", "Mozzarella"],
            "vegan": false
        },
        {
            "id": "item127",
            "name": "Mushroom Truffle Pizza",
            "cena": 17.99,
            "ingredients": ["Mozzarella", "Mushrooms", "Truffle Oil", "Arugula"],
            "vegan": true
        }
    ]`
	var menuItems []MenuItem
	err := json.Unmarshal([]byte(incomingData), &menuItems)
	if err != nil {
		fmt.Println(err)
	}

	sort.Sort(byVegan(menuItems))

	for _, mi := range menuItems {
		fmt.Printf("%+v\n", mi)
	}

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
