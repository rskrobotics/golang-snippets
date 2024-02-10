// Package mymath provides math solutions
package mymath
import "sort"


// CenteredAvg returns the average of a slice of integers,
// after removing the smallest and largest values.
func CenteredAvg(xi []int) float64{
	sort.Ints(xi)
	xi = xi[1:(len(xi)-1)]
	n :=0
	for _,v := range xi {
	n+=v
	}
	f := float64(n) / float64(len(xi))

	return f
}