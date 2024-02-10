package main

import (
	"fmt"
	"log"
	"math/rand"
)

type incorrectMathError struct {
	err      error
	severity int
}

func NewIncorrectMathError(err error) *incorrectMathError {
	e := &incorrectMathError{
		err: err,
	}
	e.assignRandomSeverity() // Assign a random severity using the method
	return e
}

func (i *incorrectMathError) assignRandomSeverity() {
	severities := []int{1, 2, 6, 8}
	i.severity = severities[rand.Intn(len(severities))]
}

func (i incorrectMathError) Error() string {
	return fmt.Sprintf("An incorrect math error occured, %w, %d", i.err, i.severity)
}

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		{
			return 0, NewIncorrectMathError(fmt.Errorf("cannot take the square root of a negative number: %f", f))
		}
	}
	return 42, nil

}
