package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output:
	// 4
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}
	tests := []test{
		{[]int{10, 20, 40, 60, 80},40},
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7 , 8, 9}, 5},
	}


	for _,v := range tests{
		f := CenteredAvg(v.data)
		if f!= v.answer {
            t.Errorf("CenteredAvg(%v) = %v, want %v", v.data, f, v.answer)
        }
	}

}
