package main

import "testing"

func TestMySum(t *testing.T) {
	sum := mySum(2, 3)
	if sum != 5 {
		t.Error(sum, "does not = 5")
	}
}
