package main

import (
	"testing"
)

func TestAssignment2(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected string
	}

	testCases := []TestCase{
		{
			Input:    "LLRR=",
			Expected: "210122",
		},
		{
			Input:    "==RLL",
			Expected: "000210",
		},
		{
			Input:    "=LLRR",
			Expected: "221012",
		},
		{
			Input:    "RRL=R",
			Expected: "012001",
		},
	}

	for _, tc := range testCases {
		actual := Solve(tc.Input)
		if actual != tc.Expected {
			t.Errorf("Case Input: %s, Expected: %s, Got: %s", tc.Input, tc.Expected, actual)
		}
	}
}
