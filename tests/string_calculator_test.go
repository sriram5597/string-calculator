package tests

import (
	"testing"

	"github.com/sriram5597/calculator/string_calculator"
	"github.com/stretchr/testify/assert"
)

func TestStringCalculator(t *testing.T) {
	testCases := []Testcase{
		{
			Name:     "Provides 0 when empty string is given",
			Input:    "",
			Expected: 0,
		},
		{
			Name:     "Provides the same number when single number is given",
			Input:    "32",
			Expected: 32,
		},
		{
			Name:     "Calculates the sum of two numbers when two numbers are given with comma separated",
			Input:    "32,43",
			Expected: 75,
		},
		{
			Name:     "Calculates the sum of all the numbers when numbers are given with comma separated",
			Input:    "32,43,25",
			Expected: 100,
		},
		{
			Name: "Calculates the sum of all the numbers when numbers are given with comma separated",
			Input: `32
			43,25`,
			Expected: 100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			res := string_calculator.Sum(testCase.Input)
			assert.Equal(t, testCase.Expected, res, "verifying result")
		})
	}
}
