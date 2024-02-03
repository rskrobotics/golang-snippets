package functions

import (
	"fmt"
)

// Define a custom type for Occupation
type Occupation int

// Define constants representing the "enum" values for Occupation
const (
    Engineer Occupation = iota
    Scientist
    Developer
)

// String returns the string representation of the Occupation
func (o Occupation) String() string {
    names := [...]string{
        "Engineer",
        "Scientist",
        "Developer",
    }

    // Prevent out-of-bounds panic
    if o < Engineer || o > Developer {
        return "Unknown"
    }

    return names[o]
}

// Worker struct with Name and Occupation
type Worker struct {
    Name        string
    Occupation  Occupation
}

// Speak function for worker types

func (w Worker) speak() {
	fmt.Println("I am ", w.Name)
}

func main() {
    // Creating an instance of Worker
    worker := Worker{"Alice", Engineer}

    // Example usage
    fmt.Printf("Worker Name: %s, Occupation: %s\n", worker.Name, worker.Occupation)

	worker.speak()
}
