package functions

import (
	"fmt"
	"log"
	"strconv"
)

type worker struct {
    name string
}

type specialist struct {
	worker
    access bool
}

func (w worker) speak() {
	fmt.Println("I am ", w.name)
}

func (s specialist) speak() {
	fmt.Println("I am ", s.name)
}

type human interface{
	speak()
}

type meaning int

func saySomething(h human){
	h.speak()
}

func (s specialist) String() string{
	return fmt.Sprint("My access has been granted: ", s.access)
}

func (m meaning) String() string{
	return fmt.Sprint("Meaning of life is ", strconv.Itoa(int(m)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log from functions_3", s.String())
}

func main()  {
	w1 := worker {
		name: "BOB ON THE JOB",
	}
	s1 := specialist {
		worker: worker{
			name: "CORN ON THE COB",
	},
	access: true,
	}

	var m meaning = 42

	saySomething(s1)
	saySomething(w1)
	fmt.Println(s1)
	fmt.Println(m)
	logInfo(s1)
	logInfo(m)
}








