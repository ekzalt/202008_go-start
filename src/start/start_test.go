package main

import "testing"

func TestSumAndMult(t *testing.T) {
	var sum int
	var mul int
	sum, mul = SumAndMult(2, 3)

	if sum != 5 || mul != 6 {
		t.Error("Expeced: 5 6, Actual:", sum, mul)
	}
}
