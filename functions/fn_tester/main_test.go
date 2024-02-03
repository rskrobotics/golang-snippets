package main

import (
	"testing"
)

func TestFactorial(t *testing.T){
	result := factorial(5)
	if result != 120{
		t.Errorf("Test incorrect, got: %d, want: %d", result, 120)
	}
}
