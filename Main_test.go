package dutch_flag_sort

import (
	"fmt"
	"math/rand"
	sort2 "sort"
	"testing"
)

func Test_correctness_sort(t *testing.T) {
	randomArray := randArray()
	want := append(randomArray[0:0:0], randomArray...)
	sort2.Ints(want)
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}},
		{[]int{2, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 2}},
		{[]int{1, 1, 1, 1, 1, 1, 1, 2}, []int{1, 1, 1, 1, 1, 1, 1, 2}},
		{[]int{2, 1, 2, 2, 1, 1, 2, 1}, []int{1, 1, 1, 1, 2, 2, 2, 2}},
		{randomArray, want},
	}

	for _, tt := range tests {
		sort(tt.input)
		if !isSliceEquals(tt.input, tt.want) {
			t.Error(fmt.Sprintf("Want %v Actual %v", tt.want, tt.input))
		}
	}
}

func x() ([]int, []int) {
	return []int{2, 1, 2, 2, 1, 1, 2, 1}, []int{1, 1, 1, 1, 2, 2, 2, 2}
}

func randArray() []int {
	ints := make([]int, rand.Intn(1000))

	for i := 0; i < cap(ints); i++ {
		ints[i] = rand.Intn(3) + 1
	}

	return ints
}

func isSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
