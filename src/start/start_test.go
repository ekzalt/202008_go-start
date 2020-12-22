package main

import "testing"

func TestSumAndMult(t *testing.T) {
	// skip unnecessary test
	// t.Skip()

	// run this case in parallel
	// t.Parallel()

	var sum int
	var mul int
	sum, mul = SumAndMult(2, 3)

	// log anything in test
	// t.Log("SumAndMult result:", sum, mul)

	if sum != 5 || mul != 6 {
		// fail this case but continue other cases
		t.Error("Expeced: 5 6, Actual:", sum, mul)

		// fail all, stop other cases
		// t.Fatal("Expeced: 5 6, Actual:", sum, mul)
	}
}
