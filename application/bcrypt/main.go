package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	s := "samplepass"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil{
		fmt.Println("ERROR FROM HASHING: ",err)
	}

	err = bcrypt.CompareHashAndPassword(bs, []byte("samplepa1ss"))
	if err != nil {
		fmt.Println("Incorrect password")
		return 
	}

	fmt.Println("Correct pass!")
}