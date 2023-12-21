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
			Name:     "Calculates the sum of two numbers when two numbers are separated by comma",
			Input:    "32,43",
			Expected: 75,
		},
		{
			Name:     "Calculates the sum of all the numbers when numbers are separated by comma",
			Input:    "32,43,25",
			Expected: 100,
		},
		{
			Name: "Calculates the sum of all the numbers when numbers are separated by comma and newline",
			Input: `32
			43,25`,
			Expected: 100,
		},
		{
			Name: "Calculates the sum of all the numbers when numbers are separated by the seperator given in the first line",
			Input: `\;
			32;43;25`,
			Expected: 100,
		},
		{
			Name: "Error should be thrown with all negative numbers when negative numbers are provided",
			Input: `\;
			32;-43;-25;-40`,
			Error: "negatives not allowed: [-43,-25,-40]",
		},
	}
	calc := string_calculator.StringCalculator{}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			res, err := calc.Sum(testCase.Input)
			if testCase.Error != "" {
				if err == nil {
					assert.Fail(t, "expecting error")
				}
				assert.Equal(t, testCase.Error, err.Error(), "verifying error message")
			} else {
				assert.Equal(t, testCase.Expected, res, "verifying result")
			}
		})
	}
	assert.Equal(t, len(testCases), calc.GetCalledCount(), "verifying the called count is equal to number test executions")
}
