package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{
			data:   []int{2, 3},
			answer: 5,
		},
		test{
			data:   []int{21, 21},
			answer: 42,
		},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Error(sum, "does not = ", v.answer)
		}
	}

}
