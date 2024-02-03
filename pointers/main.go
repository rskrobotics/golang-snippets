package main

import (
	"fmt"
	"strings"
)

type color struct {
	ral  int
	desc string
}

func changeDesc_1(c color) bool {
	c.desc = strings.ToUpper(c.desc)
	return true
}

func changeDesc_2(c *color) {
	c.desc = strings.Repeat(c.desc, 2)
}

func main() {
	black := color{
		ral:  9005,
		desc: "czarny",
	}
	white := &color{
		ral:  9003,
		desc: "biały",
	}

	fmt.Println(black.desc)
	changeDesc_1(black)
	fmt.Println(black.desc)
	changeDesc_2(&black)
	fmt.Println(black.desc)

	fmt.Println(white.desc)
	changeDesc_1(*white)
	fmt.Println(white.desc)
	changeDesc_2(white)
	fmt.Println(white.desc)

}
