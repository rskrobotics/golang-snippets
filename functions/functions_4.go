package functions

import (
	"bytes"
	// "fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	/*
		s := []byte("I am Giovanni Giorgio, but everyone calls me Giorgio")

		_, err = f.Write(s)
		if err != nil{
			log.Fatalf("error %s", err)
		}

		b := bytes.NewBufferString("APEROL")
		b.Write([]byte(" SPRITZ"))
		b.Reset()
		b.Write([]byte("BEER")) */

	// fmt.Println(b.String())

	var b bytes.Buffer
	p := person{
		first: "ADRJON",
	}

	p.writeOut(&b)
	p.writeOut(f)

}
